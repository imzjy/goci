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
	log.Println(string(body))
	log.Println("======end payload======")
}

func GitHub(w http.ResponseWriter, r *http.Request) {
	log.Println("======github payload======")
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
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
