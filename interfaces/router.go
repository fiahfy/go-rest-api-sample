package interfaces

import (
	"net/http"
	"regexp"

	"github.com/fiahfy/go-todo-rest-api/interfaces/handler"
	"github.com/fiahfy/go-todo-rest-api/interfaces/httputils"
)

type Handler func(*httputils.Context)

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type Router struct {
	Routes       []Route
	DefaultRoute Handler
}

func NewRouter(handler handler.AppHandler) *Router {
	r := &Router{
		DefaultRoute: func(ctx *httputils.Context) {
			ctx.Text(http.StatusNotFound, "Not found")
		},
	}
	r.handle(`^/hello$`, handler.GetHello)
	r.handle(`^/todos/(\d+)$`, handler.GetTodo)
	r.handle(`^/todos$`, handler.ListTodos)
	return r
}

func (a *Router) handle(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler}

	a.Routes = append(a.Routes, route)
}

func (a *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &httputils.Context{Request: r, ResponseWriter: w}

	for _, rt := range a.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	a.DefaultRoute(ctx)
}
