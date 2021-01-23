package http

import (
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"github.com/wwbweibo/EasyRoute/src/server/channel"
	"log"
	"net"
)

type HttpConnectionHandler struct {
	server *Server
}

func (handler *HttpConnectionHandler) HandleConnection(conn net.Conn) {
	if conn == nil {
		log.Fatal("connection is empty, can not handle it")
	}
	channel := channel.NewChannel(conn)
	go handler.handleRequestData(channel)
	go handler.handleResponseData(channel)
	for true {
		buffer := make([]byte, 1024)
		cnt, err := conn.Read(buffer)
		if err != nil {
			return
		}
		channel.WriteRequestData(buffer[0:cnt])
	}
}

func (handler *HttpConnectionHandler) handleRequestData(channel *channel.Channel) {
	for true {
		select {
		case <-channel.GetInputChannel():
			req := DecodeHttpRequest(channel.GetInputBuffer())
			ctx := context.NewContext(req)
			if req != nil {
				handler.server.requestDelegate(ctx)
				// write back response
				ctx.Response.Write(channel)
			}
		}
	}
}

func (handler *HttpConnectionHandler) handleResponseData(channel *channel.Channel) {
	for true {
		select {
		case <-channel.GetOutputChannel():
			buf := make([]byte, 1024)
			for true {
				cnt, err := channel.GetOutputBuffer().Read(buf)
				if err != nil || cnt == 0 {
					break
				}
				channel.GetConnection().Write(buf[0:cnt])
			}
		}

	}
}
