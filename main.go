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

	router := router.New()

	router.Get(`^/$`, h.GetIndex)
	router.Get(`^/todos/(\d+)$`, h.GetTodo)
	router.Get(`^/todos$`, h.ListTodos)
	router.Post(`^/todos$`, h.PostTodo)
	router.Put(`^/todos/(\d+)$`, h.PutTodo)
	router.Delete(`^/todos/(\d+)$`, h.DeleteTodo)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
