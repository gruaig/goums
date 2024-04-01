package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gruaig/goums/users"
	_ "github.com/lib/pq"
)

var templates map[string]*template.Template

type DisplayUsers struct {
	ID        int32
	Name      string
	Lastname  string
	Email     string
	Username  string
	Password  string
	Enabled   bool
	CreatedAt string
	UpdatedAt string
}

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["./web/content.html"] = template.Must(template.ParseFiles("./web/content.html"))
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	connectionStr := "user=postgres password=password dbname=goums sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	queries := users.New(db)

	getUsers, err := queries.ListUser(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var userList []DisplayUsers
	for _, user := range getUsers {
		u := DisplayUsers{
			ID:        user.ID,
			Name:      user.Name.String,
			Lastname:  user.Lastname.String,
			Email:     user.Email.String,
			Username:  user.Username.String,
			Password:  user.Password.String,
			Enabled:   user.Enabled.Bool,
			CreatedAt: user.CreatedAt.Time.String(),
			UpdatedAt: user.UpdatedAt.Time.String(),
		}
		userList = append(userList, u)
	}

	// Convert userList to JSON
	userListJSON, err := json.Marshal(userList)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the template with userList JSON data
	tmpl := templates["./web/content.html"]
	//tmpl.ExecuteTemplate(w, "./web/content.html", map[string]template.JS{"Users": template.JS(userListJSON)})
	tmpl.ExecuteTemplate(w, "./web/content.html", map[string]interface{}{"Users": userListJSON})

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/users", contentHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
