package main

import (
	"log"

	"github.com/gruaig/goums/sql"
)

func main() {

	connectionStr := "user=postgres password=password dbname=goums sslmode=disable"

	manager, error := db.NewConnectionManager(connectionStr)

	if err ! nil {
		log.Fatal(err)
	}



}
