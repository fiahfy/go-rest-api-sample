package main

import (
	"log"
	"net/http"

	"github.com/fiahfy/go-todo-rest-api/interfaces"
	"github.com/fiahfy/go-todo-rest-api/registry"
)

func main() {
	r := registry.NewRegistry()
	h := r.NewAppHandler()

	app := interfaces.NewRouter(h)
	err := http.ListenAndServe(":8080", app)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
