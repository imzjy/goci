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

	log.SetOutput(LogWriter{Logfile: appFullpath + ".log"})
}

type LogWriter struct {
	Logfile string
}

func (lw LogWriter) Write(p []byte) (int, error) {

	f, err := os.OpenFile(lw.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	if err != nil {
		return 0, err
	}

	n, err := f.Write(p)
	if err != nil {
		return 0, err
	}

	err = f.Sync()
	if err != nil {
		return n, err
	}

	return n, err

}
