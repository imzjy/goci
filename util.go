package main

import "strings"

func CmdSubstitute(cmd string, notify Notify) string {
	return strings.Replace(cmd, "#{branch}", notify.Branch, -1)
}
