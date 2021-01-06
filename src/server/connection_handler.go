package server

import (
	"net"
)

type IConnectionHandler interface {
	HandleConnection(conn net.Conn)
}
