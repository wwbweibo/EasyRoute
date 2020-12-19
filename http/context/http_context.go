package context

import (
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response *Response
}

func NewContext(req *http.Request) *Context {
	return &Context{
		Request:  req,
		Response: NewResponse(),
	}
}
