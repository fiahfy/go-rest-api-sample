package handler

import (
	"net/http"

	"github.com/fiahfy/go-todo-rest-api/interfaces/httputils"
)

type IndexHandler interface {
	GetIndex(ctx *httputils.Context)
}

type indexHandler struct{}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (h *indexHandler) GetIndex(ctx *httputils.Context) {
	ctx.Json(http.StatusOK, map[string]interface{}{
		"name":        "Todo API",
		"description": "Sample RESTful API With Golang",
		"homepage":    "https://github.com/fiahfy/go-todo-rest-api",
	})
}
