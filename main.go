package main

import (
	"log"
	"net/http"

	"github.com/fiahfy/go-todo-rest-api/interfaces/router"
	"github.com/fiahfy/go-todo-rest-api/registry"
)

func main() {
	r := registry.New()
	h := r.NewAppHandler()

	app := router.New(h)
	err := http.ListenAndServe(":8080", app)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
