package main

import (
	"log"
	"os"
	"path/filepath"
)

func StartLog() {

	appFullpath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(appFullpath+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)
}
