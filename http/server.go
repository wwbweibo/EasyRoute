package http

import (
	"github.com/wwbweibo/EasyRoute/server"
)

type Server struct {
	server *server.Server
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
	receiver.server.RegisterHandler(&HttpConnectionHandler{})
	return receiver.server.Serve()
}