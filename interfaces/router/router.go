package router

import (
	"net/http"
	"regexp"

	"github.com/fiahfy/go-todo-rest-api/interfaces/handler"
	"github.com/fiahfy/go-todo-rest-api/interfaces/httputils"
)

type Handler func(*httputils.Context)

type Route struct {
	Method  string
	Pattern *regexp.Regexp
	Handler Handler
}

type Router interface {
	http.Handler
}

type router struct {
	Routes       []Route
	DefaultRoute Handler
}

func New(handler handler.AppHandler) Router {
	r := &router{
		DefaultRoute: func(ctx *httputils.Context) {
			ctx.Text(http.StatusNotFound, "Not found")
		},
	}
	r.handle(http.MethodGet, `^/$`, handler.GetIndex)
	r.handle(http.MethodGet, `^/todos/(\d+)$`, handler.GetTodo)
	r.handle(http.MethodGet, `^/todos$`, handler.ListTodos)
	return r
}

func (a *router) handle(method string, pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Method: method, Pattern: re, Handler: handler}

	a.Routes = append(a.Routes, route)
}

func (a *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &httputils.Context{ResponseWriter: w, Request: r}

	for _, rt := range a.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 && rt.Method == ctx.Method {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	a.DefaultRoute(ctx)
}
