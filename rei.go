package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

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

func main() {

	// File for count record
	file := "count.conf"

	// Get last stopped transaction count
	cnt := initResume(file)
	var max uint64

	// Create new SUI client instance
	sc := new(sui.SUIClient)
	sc.Init("http://127.0.0.1:9000")

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
			for range list {

				select {

				// Graceful shutdown next transaction read
				case <-sigchan:
					cleanUp(file, cnt)

				// If ctrl-c not triggered, process the transaction
				default:

					// count always before print so we don't skip transactions
					cnt++
					fmt.Printf("Finished processing %d\n", cnt)
				}
			}
		}
	}
}
