package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type SSHCommander struct {
	User    string
	IP      string
	KeyPath string
}

func (s *SSHCommander) Command(cmd ...string) *exec.Cmd {
	arg := append(
		[]string{
			"-i",
			s.KeyPath,
			"-o",
			"StrictHostKeyChecking=no",
			fmt.Sprintf("%s@%s", s.User, s.IP),
		},
		cmd...,
	)
	return exec.Command("ssh", arg...)
}

func execSshRemote(user, host, keypath string, cmd []string) string {
	commander := SSHCommander{user, host, keypath}

	exector := commander.Command(cmd...)
	out, err := exector.Output()
	if err != nil {
		log.Panic(err)
	}
	return string(out)
}

func ExecSsh(user, host, cmd, keypath string) string {
	log.Println("ssh exec:", cmd)

	return execSshRemote(user, host, keypath, strings.Fields(cmd))
}
