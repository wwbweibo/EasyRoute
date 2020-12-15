package server

import (
	"github.com/wwbweibo/EasyRoute/server/channel"
	"github.com/wwbweibo/EasyRoute/server/http"
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {
	if conn == nil {
		log.Fatal("connection is empty, can not handle it")
	}
	channel := channel.NewChannel(conn)
	go handleRequestData(channel)
	go handleResponseData(channel)
	for true {
		buffer := make([]byte, 1024)
		cnt, err := conn.Read(buffer)
		if err != nil {
			log.Fatal("read request data error")
		}
		channel.WriteRequestData(buffer[0:cnt])
	}
}

func handleRequestData(channel *channel.Channel) {
	for true {
		select {
		case <-channel.GetInputChannel():
			http.DecodeHttpRequest(channel.GetInputBuffer())
		}
	}
}

func handleResponseData(channel *channel.Channel) {

}
