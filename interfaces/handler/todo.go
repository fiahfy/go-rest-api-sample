package handler

import (
	"net/http"
	"strconv"

	"github.com/fiahfy/go-todo-rest-api/application/usecase"
	"github.com/fiahfy/go-todo-rest-api/interfaces/httputils"
)

type TodoHandler interface {
	GetTodo(ctx *httputils.Context)
	ListTodos(ctx *httputils.Context)
}

type todoHandler struct {
	u usecase.TodoUseCase
}

func NewTodoHandler(u usecase.TodoUseCase) TodoHandler {
	return &todoHandler{u}
}

func (h *todoHandler) GetTodo(ctx *httputils.Context) {
	id, _ := strconv.Atoi(ctx.Params[0])
	todo, _ := h.u.Find(id)
	ctx.Json(http.StatusOK, todo)
}

func (h *todoHandler) ListTodos(ctx *httputils.Context) {
	todos, _ := h.u.FindAll()
	ctx.Json(http.StatusOK, todos)
}
