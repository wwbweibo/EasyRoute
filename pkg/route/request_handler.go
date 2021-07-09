package route

import (
	cctx "context"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	http2 "github.com/wwbweibo/EasyRoute/pkg/http"
	"net/http"
	"strings"
)

type httpRequestHandler struct {
	context *RouteContext
}

func (handler *httpRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := http2.HttpContext{
		Request:  request,
		Response: writer,
		Ctx:      cctx.Background(),
	}
	handler.context.app(&ctx)
}

// http请发分发

type requestHandler struct {
	routeContext *RouteContext
	delegate     delegates.RequestDelegate
}

var reqHandler = requestHandler{
	routeContext: &routeContext,
	delegate: func(ctx *http2.HttpContext) {
		request := ctx.Request
		path := strings.ToLower(request.URL.Path)
		targetNode, isMatched, err := routeContext.endPointTrie.GetMatchedRoute(path)

		if isMatched {
			endpoint := targetNode.endPoint
			if err != nil {
				ctx.Response.WriteHeader(http.StatusInternalServerError)
				ctx.Response.Write([]byte(err.Error()))
			} else if endpoint == nil || endpoint.handler == nil {
				ctx.Response.WriteHeader(http.StatusNotFound)
				ctx.Response.Write([]byte("404 Not Found"))
			} else {
				endpoint.handler(ctx)
			}
		} else {
			targetNode.defaultHandler(ctx)
		}
	},
}
