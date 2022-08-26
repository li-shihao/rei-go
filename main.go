package main

import (
	"fmt"
	"os"
)

func initResume(name string) uint64 {
	cnt := uint64(0)

	// If count file exists
	if _, err := os.Stat(name); err == nil {

		// Open count file
		file, err := os.Open(name)
		check(err)

		// Read into cnt
		_, err = fmt.Fscanf(file, "%d\n", &cnt)
		check(err)
		file.Close()
	}
	return cnt
}

func main() {

	// Get last stopped transaction count
	cnt := initResume("count.conf")

	// Create new SUI client instance
	sui := new(SUIClient)
	sui.Init("http://127.0.0.1:9000")

	// Get current total transactions count
	max := sui.GetTotalTransactionNumber()

	_, _ = sui.GetTransactionsInRange(cnt, max)

	_, _ = sui.GetTransaction("Um5bXvoCztqZlhOy/xWslobwSTrXVxVt6QxDjYG+ep0=")

}
