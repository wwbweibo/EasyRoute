package route

import (
	http2 "github.com/wwbweibo/EasyRoute/src/http"
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"net/http"
)

// http请发分发

type requestHandler struct {
	routeContext *RouteContext
	delegate     http2.RequestDelegate
}

var reqHandler = requestHandler{
	routeContext: &routeContext,
	delegate: func(ctx *context.Context) {
		request := ctx.Request
		path := request.URL.Path
		endpoint, err := routeContext.endPointTrie.GetMatchedRoute(path)

		if err != nil {
			ctx.Response.WriteHttpCode(http.StatusInternalServerError, "InternalServerError")
			ctx.Response.WriteBody([]byte(err.Error()))
		} else if endpoint == nil || endpoint.handler == nil {
			ctx.Response.WriteHttpCode(http.StatusNotFound, "NotFound")
			ctx.Response.WriteBody([]byte("404 Not Found"))
		} else {
			endpoint.handler(ctx)
		}
	},
}
