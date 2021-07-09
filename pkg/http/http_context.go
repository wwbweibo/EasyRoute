package http

import (
	"context"
	"net/http"
)

type HttpContext struct {
	Request  *http.Request
	Response http.ResponseWriter
	Ctx      context.Context
}
