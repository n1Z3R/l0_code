package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"l0/internal/models"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type Database struct {
	DB *sql.DB
}

func InitDB() *Database {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname))
	if err != nil {
		return nil
	}

	log.Println("connected!")
	return &Database{db}
}

func (d *Database) ReadById(id string) (models.Model, error) {
	var rowBytes []byte
	var order models.Model

	row := d.DB.QueryRow("select content from orders where id =$1", id)

	err := row.Scan(&rowBytes)
	if err != nil {
		return models.Model{}, err
	}
	err = json.Unmarshal(rowBytes, &order)
	if err != nil {
		return models.Model{}, err
	}
	return order, nil

}

func (d *Database) InsertModel(message []byte) (string, error) {
	var id string

	err := d.DB.QueryRow("INSERT INTO orders (content) VALUES ($1) RETURNING id", message).Scan(&id)

	if err != nil {
		return "", err
	}

	log.Println(id)
	return id, nil
}
