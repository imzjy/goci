package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var config = CiConfg{}

func BitBucket(w http.ResponseWriter, r *http.Request) {
	log.Println("======bitbucket notify======")
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read request error:", err.Error())
		return
	}

	notify, err := ParseBitBucketPayload(body)
	if err != nil {
		log.Println("parse payload error:", err.Error())
		return
	}

	trigger, err := GetMatchedTrigger(config, notify, "bitbucket")
	if err != nil {
		log.Println("no trigger for notify:", notify)
		fmt.Fprintf(w, "no trigger for notify:%#v", notify)
		return
	}

	var cmdOut []byte = []byte{}
	var cmdErr error = nil
	if trigger.Type == "local" {
		cmdOut, cmdErr = ExecLocal(trigger.Cmd, "")
	}

	if trigger.Type == "ssh" {
		cmdOut, cmdErr = ExecSsh(trigger.SshUser, trigger.SshHost, trigger.Cmd, trigger.SshKey)
	}

	if cmdErr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s", cmdErr.Error())
		log.Println("local execution error:", cmdErr.Error())
	} else {
		fmt.Fprintf(w, "%s", string(cmdOut))
		log.Println(string(cmdOut))
	}
}

func GitHub(w http.ResponseWriter, r *http.Request) {
	log.Println("======github notify======")
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read request error:", err.Error())
		return
	}

	notify, err := ParseGitHubPayload(body)
	if err != nil {
		log.Println("parse payload error:", err.Error())
		return
	}

	if notify.Ping {
		fmt.Fprintf(w, "%s", "ping ok!")
		return
	}

	trigger, err := GetMatchedTrigger(config, notify, "github")
	if err != nil {
		log.Println("no trigger for notify:", notify)
		fmt.Fprintf(w, "no trigger for notify:%#v", notify)
		return
	}

	var cmdOut []byte = []byte{}
	var cmdErr error = nil
	if trigger.Type == "local" {
		cmdOut, cmdErr = ExecLocal(trigger.Cmd, "")
	}

	if trigger.Type == "ssh" {
		cmdOut, cmdErr = ExecSsh(trigger.SshUser, trigger.SshHost, trigger.Cmd, trigger.SshKey)
	}

	if cmdErr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s", cmdErr.Error())
		log.Println("local execution error:", cmdErr.Error())
	} else {
		fmt.Fprintf(w, "%s", string(cmdOut))
		log.Println(string(cmdOut))
	}
}

func main() {

	StartLog()

	cfg, err := LoadConfig() //CAUTION!!! if variable name as config, config will totally new variable, instead of global variable
	if err != nil {
		log.Fatal(err)
	}
	config = *cfg

	http.HandleFunc("/bitbucket", BitBucket)
	http.HandleFunc("/github", GitHub)

	log.Println("goci start to listening on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
