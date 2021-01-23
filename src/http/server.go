package http

import (
	"github.com/wwbweibo/EasyRoute/src/server"
)

type Server struct {
	server          *server.Server
	requestDelegate RequestDelegate
}

func NewHttpServer(host, port string, delegate RequestDelegate) *Server {
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "80"
	}
	return &Server{
		requestDelegate: delegate,
		server:          server.NewServer(host, port),
	}
}

func (receiver *Server) Serve() error {
	receiver.server.RegisterHandler(&HttpConnectionHandler{
		server: receiver,
	})
	return receiver.server.Serve()
}
