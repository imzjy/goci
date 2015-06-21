package main

import (
	"log"
	"os/exec"
	"strings"
)

func execLocalCmd(cmd, dir string) ([]byte, error) {
	log.Println("command is ", cmd)
	// splitting to head(cmd) and parts(args)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	exector := exec.Command(head, parts...)
	exector.Dir = dir
	out, err := exector.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func ExecLocal(script, dir string) ([]byte, error) {
	out, err := execLocalCmd(script, dir)

	return out, err
}
