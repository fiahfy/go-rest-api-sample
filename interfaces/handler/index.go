package handler

import (
	"net/http"

	"github.com/fiahfy/go-todo-rest-api/interfaces/router"
)

type IndexHandler interface {
	GetIndex(c *router.Context)
}

type indexHandler struct{}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (h *indexHandler) GetIndex(c *router.Context) {
	c.Json(http.StatusOK, map[string]interface{}{
		"name":        "Todo API",
		"description": "Sample RESTful API With Golang",
		"homepage":    "https://github.com/fiahfy/go-todo-rest-api",
	})
}
