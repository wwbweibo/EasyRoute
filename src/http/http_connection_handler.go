package http

import (
	"fmt"
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
				fmt.Println("Ready to process request")
				handler.server.routes.HandleRequest(ctx)
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
			cnt, _ := channel.GetOutputBuffer().Read(buf)
			channel.GetConnection().Write(buf[0:cnt])
		}

	}
}
