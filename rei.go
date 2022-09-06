package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
	"fmt"

	"rei.io/rei/internal/database"
	"rei.io/rei/internal/helpers"
	"rei.io/rei/internal/sui"
	"rei.io/rei/server"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
)

var check = helpers.Check

func processTX(thread chan int, transactionId string, sc *sui.SUIClient, db *database.EntClient, cnt uint64, firstLoadLimit uint64) {

	tx, err := sc.GetTransaction(transactionId)
	check(err)
	db.CreateTransaction(tx)

	// If call has arguments insert arguments
	if tx.Arguments != nil && tx.GetStatus() {
		for _, k := range *tx.Arguments {
			db.CreateArgument(k)
		}
	}

	if tx.GetType() == "Publish" && tx.GetStatus() {
		dply, err := tx.GetContractDeploy()
		check(err)
		db.CreatePackage(dply)
	}

	// Insert events
	if tx.Events != nil {
		for _, k := range *tx.Events {

			// Insert event
			db.CreateEvent(k)

			// First 78k dont need multiple insertions for accounts
			if cnt > firstLoadLimit || !db.QueryAccountFirstLoad(k.Sender, firstLoadLimit) {
				// Insert sender account
				sdr, err := sc.GetAccount(k.Sender)
				check(err)
				db.CreateAccount(sdr, cnt)

				// Insert sender NFTs

				if sdr.GetAccountNFTs() != nil {
					for _, k := range sdr.GetAccountNFTs() {
						db.CreateNFT(k, cnt)
					}
				}
			}

			if cnt > firstLoadLimit || !db.QueryAccountFirstLoad(k.Sender, firstLoadLimit) {
				// Check if recipient exist && recipient is actually an address (not 'shared' or smth else)
				if k.Recipient != nil && strings.HasPrefix(*k.Recipient, "0x") {
					rcp, err := sc.GetAccount(*k.Recipient)
					check(err)
					db.CreateAccount(rcp, cnt)

					if rcp.GetAccountNFTs() != nil {
						for _, k := range rcp.GetAccountNFTs() {
							db.CreateNFT(k, cnt)
						}
					}
				}
			}

			if cnt > firstLoadLimit || !db.QueryObjectFirstLoad(k.ObjectId, firstLoadLimit) {
				obj, err := sc.GetObject(k.ObjectId)
				check(err)
				db.CreateObject(obj, cnt)
			}
		}
	}

	// count always before print so we don't skip transactions
	log.Printf("%s: Finished processing %d\n", tx.GetID(), cnt)
	<-thread
}

func main() {

	// File for count record
	file := "count.conf"

	// Get last stopped transaction count
	cnt := helpers.InitResume(file)
	var max uint64

	// Create new SUI client instance
	sc := new(sui.SUIClient)
	sc.Init("http://178.20.44.135:9000")

	// New db instance
	db := new(database.EntClient)
	//connStr := "host=localhost port=5432 user=postgres dbname=rei password=postgres sslmode=disable"

	p := postgres.Preset(
		postgres.WithUser("gnomock", "gnomick"),
		postgres.WithDatabase("mydb"),
	)

	log.Println("Starting docker....")
	container, _ := gnomock.Start(p)

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(),
		"gnomock", "gnomick", "mydb",
	)

	db.Init("postgres", connStr)

	// API server set-up
	r := server.CreateServer(connStr)

	go func() {
		http.ListenAndServe(":6060", r)
	}()

	// Listener to kill
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Goroutine settings
	const MAX = 1
	thread := make(chan int, MAX)

	// Only used as a limit
	firstLoadLimit := sc.GetTotalTransactionNumber()

	// Main loop
	for {
		max = sc.GetTotalTransactionNumber()

		/*
			Only triggered when max has been reached
			Loop will run until max increases
		*/
		if max == cnt {
			// This loop runs every 10 seconds
			select {
			case <-sigchan:
				helpers.CleanUp(file, cnt)

			// This basically allows us to wait for signal for 10 seconds
			case <-time.After(10 * time.Second):
			}

		} else {

			/*
				Calculate how many transactions we need to retreve this iteration
				If the diff is less than 4096, retrieve the diff
				Otherwise retrieve 4096
			*/
			add := (func() uint64 {
				if max-cnt > 4096 {
					return 4096
				} else {
					return max - cnt
				}
			})

			list, err := sc.GetTransactionsInRange(cnt, cnt+add())
			check(err)

			// We are listening for signal during *every transaction read*
			for _, v := range list {

				select {

				// Graceful shutdown next transaction read
				case <-sigchan:
					helpers.CleanUp(file, cnt)

				// If ctrl-c not triggered, process the transaction
				default:
					thread <- 1
					cnt++
					go processTX(thread, v, sc, db, cnt, firstLoadLimit)
				}
			}

			// Intermediate saving
			helpers.Save(file, cnt)
		}
	}
}
