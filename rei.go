package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	database "rei.io/rei/internal/db"
	"rei.io/rei/internal/helpers"
	"rei.io/rei/internal/sui"
)

var check = helpers.Check

func initResume(name string) uint64 {
	cnt := uint64(0)

	// If count file exists
	if _, err := os.Stat(name); err == nil {

		// Open count file
		file, err := os.Open(name)
		check(err)

		// Read into cnt
		_, err = fmt.Fscanf(file, "%d", &cnt)
		check(err)
		file.Close()
	}
	return cnt
}

func cleanUp(name string, cnt uint64) {

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	check(err)

	err = file.Truncate(0)
	check(err)

	_, err = file.Seek(0, 0)
	check(err)

	_, err = fmt.Fprintf(file, "%d", cnt)
	check(err)

	fmt.Println("Program killed !")
	os.Exit(0)
}

func processTX(transactionId string, sc *sui.SUIClient, db *database.EntClient, cnt uint64) {
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

			// Insert sender account
			sdr, err := sc.GetAccount(k.Sender)
			check(err)
			db.CreateAccount(sdr)

			// Insert sender NFTs
			if sdr.GetAccountNFTs() != nil {
				for _, k := range sdr.GetAccountNFTs() {
					db.CreateNFT(k)
				}
			}

			// Check if recipient exist && recipient is actually an address (not 'shared' or smth else)
			if k.Recipient != nil && strings.HasPrefix(*k.Recipient, "0x") {
				rcp, err := sc.GetAccount(*k.Recipient)
				check(err)
				db.CreateAccount(rcp)

				if rcp.GetAccountNFTs() != nil {
					for _, k := range rcp.GetAccountNFTs() {
						db.CreateNFT(k)
					}
				}
			}

			obj, err := sc.GetObject(k.ObjectId)
			check(err)
			db.CreateObject(obj)
		}
	}

	// count always before print so we don't skip transactions
	fmt.Printf("%s: Finished processing %d\n", tx.GetID(), cnt)
}

func main() {

	// File for count record
	file := "count.conf"

	// Get last stopped transaction count
	cnt := initResume(file)
	var max uint64

	// Create new SUI client instance
	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

	db := new(database.EntClient)
	db.Init("postgres", "host=localhost port=5432 user=postgres dbname=rei password=postgres sslmode=disable")

	// Listener to kill
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Mimic a do while loop... Always execute at least once
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
				cleanUp(file, cnt)

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
					cleanUp(file, cnt)

				// If ctrl-c not triggered, process the transaction
				default:
					cnt++
					processTX(v, sc, db, cnt)
				}
			}
		}
	}
}
