package http

import (
	"net/http"
)

type HttpContext struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func NewContext(req *http.Request, response http.ResponseWriter) *HttpContext {
	return &HttpContext{
		Request:  req,
		Response: response,
	}
}
