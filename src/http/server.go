package http

import (
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"github.com/wwbweibo/EasyRoute/src/server"
)

type Server struct {
	server *server.Server
	routes *route.RouteContext
}

func NewHttpServer(host, port string) *Server {

	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "80"
	}
	return &Server{
		server: server.NewServer(host, port),
	}
}

func (receiver *Server) Serve() error {
	receiver.server.RegisterHandler(&HttpConnectionHandler{
		server: receiver,
	})
	return receiver.server.Serve()
}

func (receiver *Server) RegisterHandlers(ctx *route.RouteContext) {
	receiver.routes = ctx
	receiver.routes.InitRoute()
}
