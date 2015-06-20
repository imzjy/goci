package main

import "strings"

func CmdSubstitute(cmd string, notify Notify) string {
	newCmd := strings.Replace(cmd, "#{branch}", notify.Branch, -1)
	return strings.Replace(newCmd, "#{repo}", notify.Repository, -1)
}
