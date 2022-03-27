package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "user"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Server started")

	// connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s  sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Successfully connected:", db.Ping() == nil)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello Thomas! How are you today ?", "response":"I'm fine."}`))
}
