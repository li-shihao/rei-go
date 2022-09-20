package helpers

import (
	"fmt"
	"os"
)

var check = Check

// Finds a number in a given file; if file does not exist assume the number to be 0
func InitResume(name string) int64 {
	cnt := int64(0)

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

// Truncates and saves a number into a given file
func Save(name string, cnt int64) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	check(err)

	err = file.Truncate(0)
	check(err)

	_, err = file.Seek(0, 0)
	check(err)

	_, err = fmt.Fprintf(file, "%d", cnt)
	check(err)
}

// Truncates and saves a number into a given file, then exits
func CleanUp(name string, cnt int64) {

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
