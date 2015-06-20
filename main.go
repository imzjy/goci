package main

import (
	"fmt"
	"log"
	"net/http"
)

var config = CiConfg{}

func BitBucket(w http.ResponseWriter, r *http.Request) {
	log.Println("======bitbucket payload======")
	log.Println(string(r.Body))
	log.Println("======end payload======")
}

func GitHub(w http.ResponseWriter, r *http.Request) {
	log.Println("======github payload======")
	log.Println(string(r.Body))
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
	log.Fatal(http.ListenAndServe(":" + config.Port, nil))
}
