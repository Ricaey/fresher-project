package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"fresher-project/rear/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/comment/get", GetComments)
	http.HandleFunc("/comment/add", AddComment)
	http.HandleFunc("/comment/delete", DeleteComment)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong~")
	})
	log.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
