package delegate

import (
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"net/http"
)

var NotFoundDelegate = func(ctx *context.Context) {
	ctx.Response.WriteHttpCode(http.StatusNotFound, "NotFound")
	ctx.Response.WriteBody([]byte("404 Not Found"))
}
