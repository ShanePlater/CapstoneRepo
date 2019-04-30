package orm

import (
	"log"
	"os"
)

// checkErr handles the error.
func checkErr(err error) {
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
