package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connectionStr := "user=postgres password=password dbname=goums sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the database connection when main function exits

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// ctx := context.Background()
	// queries := users.New(db)

	// getUsers, err := queries.ListUser(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Iterate over the results
	// for _, user := range getUsers {
	// 	fmt.Println("User:", user.Name.String)
	// }

}
