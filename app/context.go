package app

import (
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (c *Context) Text(code int, body string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.WriteHeader(code)

	io.WriteString(c.ResponseWriter, fmt.Sprintf("%s\n", body))
}
