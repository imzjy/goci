package main

import (
	"log"
	"os"
	"path/filepath"
)

func StartLog() {
	logFilename := ""

	appFullpath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}

	logFilename = appFullpath + ".log"
	logfile, err := os.Open(logFilename)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logfile)
}
