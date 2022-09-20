package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/nxadm/tail"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"rei.io/rei/internal/database"
	"rei.io/rei/internal/helpers"
	"rei.io/rei/internal/sui"
	"rei.io/rei/server"
)

var check = helpers.Check
var cache = make(map[string]bool)
var cacheStore = make(map[string]bool)
var c sync.Mutex
var cu sync.Mutex

func readCache(filePath string, db *database.EntClient, curLine *int) {

	t, err := tail.TailFile(
		"/home/arthur/sui/cache", tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		data := []byte(line.Text)

		tempMap := make(map[string]interface{})
		cache := make(map[string]interface{})
		json.Unmarshal(data, &tempMap)
		json.Unmarshal(data, &cache)

		delete(cache, "transaction_id")

		// Make sure duplicate dont get in
		cu.Lock()
		notWritten := !cacheStore[fmt.Sprintf("%+v", cache)]
		cu.Unlock()
		if notWritten {
			cu.Lock()
			cacheStore[fmt.Sprintf("%+v", cache)] = true
			cu.Unlock()

			tempMap = helpers.SocketStrip(tempMap)

			var tempObj sui.SocketObj
			err = mapstructure.Decode(tempMap, &tempObj)
			helpers.Check(err)

			db.CreateObjectFromSocket(tempObj)
			log.Printf("%s: Finished processing object", tempObj.Id)
		}

		*curLine = line.Num
	}
}

func processTX(thread chan int, transactionId string, sc *sui.SUIClient, db *database.EntClient, cnt int64, firstLoadLimit int64) {

	tx, err := sc.GetTransaction(transactionId)
	check(err)

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

			sender := k.Sender
			recipient := k.Recipient

			// Haven't accessed yet

			// First 78k dont need multiple insertions for accounts
			if cnt > firstLoadLimit {

				// Insert sender account
				sdr, err := sc.GetAccount(sender)
				if err != nil {
					db.CreateDeletedAccount(sender, cnt)
				} else {
					db.UpsertAccount(sdr, cnt)
					db.UpdateAccount(sender, cnt, tx.GetID())

					// Insert sender NFTs
					if sdr.GetAccountNFTs() != nil {
						for _, k := range sdr.GetAccountNFTs() {
							db.CreateNFT(k, cnt)
						}
					}
				}

			} else {
				c.Lock()
				condition := cache[sender]
				c.Unlock()
				if !condition {
					c.Lock()
					if cache[sender] {
						panic(1)
					}
					cache[sender] = true
					c.Unlock()
					// Insert sender account
					sdr, err := sc.GetAccount(sender)
					if err != nil {
						db.CreateDeletedAccount(sender, cnt)
					} else {
						db.UpsertAccount(sdr, cnt)
						db.UpdateAccount(sender, cnt, tx.GetID())

						// Insert sender NFTs
						if sdr.GetAccountNFTs() != nil {
							for _, k := range sdr.GetAccountNFTs() {
								db.CreateNFT(k, cnt)
							}
						}
					}
				} else {
					db.UpdateAccount(sender, cnt, tx.GetID())
				}
			}

			if recipient != nil && strings.HasPrefix(*recipient, "0x") {

				if cnt > firstLoadLimit {
					// Check if recipient exist && recipient is actually an address (not 'shared' or smth else)
					rcp, err := sc.GetAccount(*recipient)
					if err != nil {
						db.CreateDeletedAccount(*recipient, cnt)
					} else {
						db.UpsertAccount(rcp, cnt)
						db.UpdateAccount(*recipient, cnt, tx.GetID())

						// Insert sender NFTs
						if rcp.GetAccountNFTs() != nil {
							for _, k := range rcp.GetAccountNFTs() {
								db.CreateNFT(k, cnt)
							}
						}
					}
				} else {
					c.Lock()
					condition := cache[*recipient]
					c.Unlock()
					if !condition {
						c.Lock()
						if cache[*recipient] {
							panic(1)
						}
						cache[*recipient] = true
						c.Unlock()
						// Check if recipient exist && recipient is actually an address (not 'shared' or smth else)
						rcp, err := sc.GetAccount(*recipient)
						if err != nil {
							db.CreateDeletedAccount(*recipient, cnt)
						} else {
							db.UpsertAccount(rcp, cnt)
							db.UpdateAccount(*recipient, cnt, tx.GetID())

							// Insert sender NFTs
							if rcp.GetAccountNFTs() != nil {
								for _, k := range rcp.GetAccountNFTs() {
									db.CreateNFT(k, cnt)
								}
							}
						}
					} else {
						db.UpdateAccount(*recipient, cnt, tx.GetID())
					}
				}
			}
		}
	}
	db.CreateTransaction(tx)

	// count always before print so we don't skip transactions
	log.Printf("%s: Finished processing transaction %d\n", tx.GetID(), cnt)
	<-thread
}
func main() {

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

	log.Println(connStr)

	var lineCur int
	db.FirstRun("postgres", connStr)

	go readCache("~/sui/cache", db, &lineCur)

	// API server set-up
	r := server.CreateServer(connStr)

	go func() {
		http.ListenAndServe(":6060", r)
	}()

	// File for count record
	file := "count.ini"
	lineFile := "line.ini"

	// Get last stopped transaction count
	cnt := helpers.InitResume(file)
	var max int64

	// Create new SUI client instance
	sc := new(sui.SUIClient)
	block := sc.Init("http://127.0.0.1:9000")
	<-block

	// Listener to kill
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Goroutine settings
	const MAX = 4
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
				helpers.Save(lineFile, cnt)
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
			add := (func() int64 {
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
			helpers.Save(lineFile, cnt)
		}
	}
}
