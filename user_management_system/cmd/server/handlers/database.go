package handlers

import (
	"database/sql"
	"log"
)

func pingTest() {
	connecStr := "postgres://postgres:password@localhost:5432?goums?sslmode=disable"
	db, err := sql.Open("postgres", connecStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
