package route

import (
	cctx "context"
	delegates2 "github.com/wwbweibo/EasyRoute/delegates"
	http3 "github.com/wwbweibo/EasyRoute/http"
	"github.com/wwbweibo/EasyRoute/log"
	"net/http"
	"strings"
)

type HttpRequestHandler struct {
	RequestDelegate delegates2.RequestDelegate
}

func (handler *HttpRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := http3.Context{
		Request:  request,
		Response: writer,
		Ctx:      cctx.Background(),
	}
	handler.RequestDelegate(&ctx)
}

// http请发分发

type requestHandler struct {
	delegate delegates2.RequestDelegate
}

func newRequestHandler(trie *EndPointTrie) requestHandler {
	return requestHandler{
		delegate: func(ctx *http3.Context) {
			request := ctx.Request
			path := strings.ToLower(request.URL.Path)
			log.Info("[RequestHandler] request for path " + path)
			targetNode, isMatched, err := trie.GetMatchedRoute(path)

			if isMatched {
				endpoint := targetNode.endPoint
				if err != nil {
					log.Error("[RequestHandler] error while searching route", err)
					ctx.Response.WriteHeader(http.StatusInternalServerError)
					ctx.Response.Write([]byte(err.Error()))
				} else if endpoint == nil || endpoint.handler == nil {
					log.Info("[RequestHandler] route matched but no route handler find for path " + path)
					ctx.Response.WriteHeader(http.StatusNotFound)
					ctx.Response.Write([]byte("404 Not Found"))
				} else {
					endpoint.handler(ctx)
				}
			} else {
				log.Info("[RequestHandler] no route match for " + path + " execute default handler")
				targetNode.DefaultHandler(ctx)
			}
		},
	}
}
