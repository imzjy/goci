package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
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

var writeMutex sync.Mutex

func (lw LogWriter) Write(p []byte) (int, error) {
	writeMutex.Lock()

	f, err := os.OpenFile(lw.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := f.Write(p)
	if err != nil {
		return 0, err
	}

	err = f.Sync()
	if err != nil {
		return n, err
	}

	stat, err := f.Stat()
	if err != nil {
		return n, err
	}

	//log rotate
	var extension = filepath.Ext(lw.Logfile)
	var nameWithoutExt = lw.Logfile[0 : len(lw.Logfile)-len(extension)]
	megaBytes := 1024 * 1024
	if stat.Size() > int64(10*megaBytes) {
		err := os.Rename(lw.Logfile, nameWithoutExt+"."+time.Now().UTC().Format("2006-01-02T15:04:05Z")+".log")
		if err != nil {
			return n, err
		}
	}

	writeMutex.Unlock()
	return n, err
}
