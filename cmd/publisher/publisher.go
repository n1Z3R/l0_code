package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sendFile("../../model.json")
	time.Sleep(5 * time.Second)
	sendFile("../../invalid_model.json")

}
func sendFile(path string) {
	file, err := ioutil.ReadFile(path)
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
