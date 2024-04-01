package main

import (
	"log"

	"github.com/gruaig/goums/sql"
)

func main() {
	connectionStr := "user=postgres password=password dbname=goums sslmode=disable"

	// Create a new connection manager
	manager, err := sql.NewConnectionManager(connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	defer manager.Close() // Close the connection when main() exits

	err = manager.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
