package main

import (
	"database/sql/driver"
	"log"
	"net/http"
	"os"

	"github.com/Encrypto07/car-app/datastore/car"
)

func main() {
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}
	var err error

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}
	datastore := car.New(db)
	handler := handler.Car.New(datastore)

	http.HandleFunc("/car", handler.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
