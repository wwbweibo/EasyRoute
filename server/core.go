package server

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	address  string
	port     string
	signal   sync.Mutex
	connChan chan net.Conn
}

func NewServer(host, port string) *Server {
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "80"
	}
	return &Server{
		address: host,
		port:    port,
	}
}

func (receiver *Server) Serve() error {
	server, err := net.Listen("tcp", receiver.address+":"+receiver.port)
	if err != nil {
		return err
	}
	defer server.Close()
	for true {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("error to accept request")
		}
		go HandleConnection(conn)
	}
	return nil
}
