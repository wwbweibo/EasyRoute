package route

import (
	cctx "context"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	http2 "github.com/wwbweibo/EasyRoute/pkg/http"
	"net/http"
	"strings"
)

type HttpRequestHandler struct {
	RequestDelegate delegates.RequestDelegate
}

func (handler *HttpRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := http2.HttpContext{
		Request:  request,
		Response: writer,
		Ctx:      cctx.Background(),
	}
	handler.RequestDelegate(&ctx)
}

// http请发分发

type requestHandler struct {
	delegate delegates.RequestDelegate
}

func newRequestHandler(trie *EndPointTrie) requestHandler {
	return requestHandler{
		delegate: func(ctx *http2.HttpContext) {
			request := ctx.Request
			path := strings.ToLower(request.URL.Path)
			targetNode, isMatched, err := trie.GetMatchedRoute(path)

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
				targetNode.DefaultHandler(ctx)
			}
		},
	}
}
