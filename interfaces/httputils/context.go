package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (c *Context) Text(code int, body string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.WriteHeader(code)
	fmt.Fprintf(c.ResponseWriter, fmt.Sprintf("%s\n", body))
}

func (c *Context) Json(code int, data interface{}) {
	o := &Response{code, data}
	res, err := json.Marshal(o)
	if err != nil {
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.WriteHeader(code)
	c.ResponseWriter.Write(res)
}
