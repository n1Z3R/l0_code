package main

import (
	"encoding/json"
	"l0/internal/app"
	"l0/internal/cache"
	"l0/internal/database"
	"l0/internal/models"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/nats-io/stan.go"
)

const (
	address = "127.0.0.1:8081"
)

func main() {
	sc, err := stan.Connect("test-cluster", "client")
	defer func(sc stan.Conn) {
		if err = sc.Close(); err != nil {
			log.Println(err)
		}
	}(sc)

	mainApp := app.App{Database: *database.InitDB(), Cache: *cache.NewCache()}
	defer mainApp.DB.Close()
	sc.Subscribe("foo", func(msg *stan.Msg) {
		var order models.Model
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println(err)
			return
		}

		var valid models.Model
		if reflect.DeepEqual(order, valid) {
			log.Println("invalid model")
			return
		}

		id, err := mainApp.Database.InsertModel(msg.Data)
		if err != nil {
			log.Println(err)
			return
		}
		mainApp.Cache.Add(id, order)
	}, stan.StartAtTimeDelta(30*time.Second))

	server := &http.Server{
		Addr:    address,
		Handler: mainApp.GetHandler(),
	}

	server.ListenAndServe()

}
