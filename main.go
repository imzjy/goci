package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var config = CiConfg{}

func BitBucket(w http.ResponseWriter, r *http.Request) {
	log.Println("======bitbucket notify======")

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	notify, _ := ParseBitBucketPayload(body)

	trigger, err := GetMatchedTrigger(config, notify, "bitbucket")
	if err != nil {
		log.Println("no trigger for notify:", notify)
		return
	}

	if trigger.Type == "local" {
		log.Println(ExecLocal(trigger.Cmd, ""))
	}
	if trigger.Type == "ssh" {
		log.Println(ExecSsh(trigger.SshUser, trigger.SshHost, trigger.Cmd, trigger.SshKey))
	}
}

func GitHub(w http.ResponseWriter, r *http.Request) {
	log.Println("======github notify======")

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	notify, _ := ParseGitHubPayload(body)

	trigger, err := GetMatchedTrigger(config, notify, "github")
	if err != nil {
		log.Println("no trigger for notify:", notify)
		return
	}

	if trigger.Type == "local" {
		log.Println(ExecLocal(trigger.Cmd, ""))
	}
	if trigger.Type == "ssh" {
		log.Println(ExecSsh(trigger.SshUser, trigger.SshHost, trigger.Cmd, trigger.SshKey))
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
