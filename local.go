package main

import (
	"log"
	"os/exec"
)

func execLocalCmd(cmd string) ([]byte, error) {
	log.Println("local cmd:", cmd)

	exector := exec.Command("/bin/sh", "-c", cmd)
	out, err := exector.Output()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}

func ExecLocal(script string) ([]byte, error) {
	return execLocalCmd(script)
}
