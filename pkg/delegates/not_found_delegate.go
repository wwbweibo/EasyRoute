package delegates

import (
	http2 "github.com/wwbweibo/EasyRoute/pkg/http"
	"net/http"
)

var NotFoundDelegate = func(ctx *http2.HttpContext) {
	ctx.Response.WriteHeader(http.StatusNotFound)
	ctx.Response.Write([]byte("404 Not Found"))
}
