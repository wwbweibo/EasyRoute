package http

import (
	"context"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Ctx      context.Context
}
