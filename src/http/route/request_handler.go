package route

import (
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"net/http"
)

// http请发分发

type requestHandler struct {
	routeContext *RouteContext
	delegate     RequestDelegate
}

var reqHandler = requestHandler{
	routeContext: &routeContext,
	delegate: func(ctx *context.Context) {
		request := ctx.Request
		path := request.URL.Path
		endpoint, err := routeContext.endPointTrie.GetMatchedRoute(path)

		if err != nil {
			ctx.Response.WriteHttpCode(http.StatusInternalServerError)
			ctx.Response.WriteBody([]byte(err.Error()))
		} else if endpoint == nil || endpoint.handler == nil {
			ctx.Response.WriteHttpCode(http.StatusNotFound)
			ctx.Response.WriteBody([]byte("404 Not Found"))
		} else {
			endpoint.handler(ctx)
		}
	},
}
