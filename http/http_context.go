package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Ctx      context.Context
}

func (context *Context) Write(content []byte, status int, contentType string) error {
	context.Response.WriteHeader(status)
	_, err := context.Response.Write(content)
	context.Request.Header.Add("Content-Type", contentType)
	return err
}

func (context *Context) WriteJson(content interface{}, status int) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}
	return context.Write(data, status, "application/json")
}

func (context *Context) WritePlainText(content string, status int) error {
	return context.Write([]byte(content), status, "text/plain")
}
