package handler

import (
	"net/http"

	"github.com/fiahfy/go-todo-rest-api/interfaces/httputils"
)

type HelloHandler interface {
	GetHello(ctx *httputils.Context)
}

type helloHandler struct{}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) GetHello(ctx *httputils.Context) {
	ctx.Text(http.StatusOK, "Hello world")
}
