package main

import (
	"log"
	"net/http"

	"github.com/fiahfy/go-todo-rest-api/interfaces"
)

func main() {
	r := interfaces.NewRouter()

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
