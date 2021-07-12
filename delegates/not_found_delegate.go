package delegates

import (
	http3 "github.com/wwbweibo/EasyRoute/http"
	"net/http"
)

var NotFoundDelegate = func(ctx *http3.HttpContext) {
	ctx.Response.WriteHeader(http.StatusNotFound)
	ctx.Response.Write([]byte("404 Not Found"))
}
