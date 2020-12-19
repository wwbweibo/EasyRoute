package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	address  string
	port     string
	signal   sync.Mutex
	connChan chan net.Conn
	handler  IConnectionHandler
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
	if receiver.handler == nil {
		panic("serve error, no handler registered")
	}
	server, err := net.Listen("tcp", receiver.address+":"+receiver.port)
	if err != nil {
		return err
	}
	defer server.Close()
	fmt.Printf("Listen on address %s:%s\n", receiver.address, receiver.port)
	for true {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("error to accept request")
		}
		go receiver.handler.HandleConnection(conn)
	}
	return nil
}

func (receiver *Server) RegisterHandler(handler IConnectionHandler) {
	receiver.handler = handler
}
