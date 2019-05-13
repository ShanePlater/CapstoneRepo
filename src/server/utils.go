package main

import (
	"flag"
	"log"
	"os"
)

// Get configuration file directory.
func getDir() string {
	// If configuration file directory is set via command line flag.
	if *flag.Bool("config", false, "config file path") {
		flag.Parse()
		return flag.Args()[0]
	}

	// If configuration file directory is not set, it is assumed that
	// configuration file is located in same directory as this application.
	dir, err := os.Getwd()
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	return dir
}
