package http

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/server/channel"
	"log"
	"net"
)

type HttpConnectionHandler struct {
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
			ctx := Context{}
			req := DecodeHttpRequest(channel.GetInputBuffer())
			if req != nil {
				ctx.Request = req
				fmt.Println("Ready to process request")
			}

			// todo : dispatch request
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
