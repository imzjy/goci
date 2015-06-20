package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var config = CiConfg{}

func BitBucket(w http.ResponseWriter, r *http.Request) {
	log.Println("======bitbucket payload======")
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	notify, _ := ParseBitBucketPayload(body)
	log.Println(notify)

	trigger, err := GetMatchedTrigger(config, notify, "bitbucket")
	if err != nil {
		//ignore
		log.Println("no trigger for notify:", notify)
	}

	if trigger.Type == "local" {
		ExecLocal(trigger.Cmd, "")
	}
	if trigger.Type == "ssh" {
		ExecSsh(trigger.SshUser, trigger.SshHost, trigger.Cmd, trigger.SshKey)
	}
	log.Println("======end payload======")
}

func GitHub(w http.ResponseWriter, r *http.Request) {
	log.Println("======github payload======")
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	notify, _ := ParseGitHubPayload(body)
	log.Println(notify)

	trigger, err := GetMatchedTrigger(config, notify, "github")
	if err != nil {
		//ignore
		log.Println("no trigger for notify:", notify)
	}

	if trigger.Type == "local" {
		ExecLocal(trigger.Cmd, "")
	}
	if trigger.Type == "ssh" {
		ExecSsh(trigger.SshUser, trigger.SshHost, trigger.Cmd, trigger.SshKey)
	}

	log.Println("======end payload======")
}

func main() {

	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)

	http.HandleFunc("/bitbucket", BitBucket)
	http.HandleFunc("/github", GitHub)

	log.Println("goci start to listening on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
