package main

import (
	"go-todo-rest-api/app"
	"log"
	"net/http"
)

func main() {
	app.InitStore()

	r := app.NewRouter()

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
