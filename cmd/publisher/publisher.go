package main

import (
	"io/ioutil"
	"log"

	"github.com/nats-io/stan.go"
)

func main() {
	file, err := ioutil.ReadFile("../../not_model.json")
	if err != nil {
		log.Println(err)
		return
	}
	sc, err := stan.Connect("test-cluster", "main")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(sc stan.Conn) {
		if err = sc.Close(); err != nil {
		}
	}(sc)

	err = sc.Publish("foo", file)
	if err != nil {
		log.Fatalln(err)
	}

}
