package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:8080")

	http.HandleFunc("/", list)
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list")
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create")
}

func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "read")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete")
}
