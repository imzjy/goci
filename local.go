package main

import (
	"log"
	"os/exec"
	"strings"
)

func execLocalCmd(cmd, dir string) []byte {
	log.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	exector := exec.Command(head, parts...)
	exector.Dir = dir
	out, err := exector.Output()
	if err != nil {
		log.Panic(err)
	}
	return out
}

func ExecLocal(script, dir string) string {
	out := execLocalCmd(script, dir)

	return string(out)
}
