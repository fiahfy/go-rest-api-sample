package handler

import (
	"net/http"
	"strconv"

	"github.com/fiahfy/go-todo-rest-api/application/usecase"
	"github.com/fiahfy/go-todo-rest-api/interfaces/router"
)

type TodoHandler interface {
	GetTodo(c *router.Context)
	ListTodos(c *router.Context)
}

type todoHandler struct {
	u usecase.TodoUseCase
}

func NewTodoHandler(u usecase.TodoUseCase) TodoHandler {
	return &todoHandler{u}
}

func (h *todoHandler) GetTodo(c *router.Context) {
	id, _ := strconv.Atoi(c.Params[0])
	todo, _ := h.u.Find(id)
	c.Json(http.StatusOK, todo)
}

func (h *todoHandler) ListTodos(c *router.Context) {
	todos, _ := h.u.FindAll()
	c.Json(http.StatusOK, todos)
}
