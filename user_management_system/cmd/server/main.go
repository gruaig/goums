package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/gruaig/goums/users"
	_ "github.com/lib/pq"
)

func main() {

	connecStr := "postgres://postgres:password@localhost:5432?goums?sslmode=disable"
	db, err := sql.Open("postgres", connecStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ctx := context.Background()
	queries := users.New(db)

}
