package server

import (
	"fmt"
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {
	if conn == nil {
		log.Fatal("connection is empty, can not handle it")
	}
	channel := NewChannel(conn)
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

func handleRequestData(channel *Channel) {
	for true {
		select {
		case <-channel.inputChan:
			buffer := make([]byte, 1024)
			cnt, _ := channel.inputBuffer.Read(buffer)
			fmt.Println(string(buffer[:cnt]))
			// todo : invoke decoder, and handler
		}
	}
}

func handleResponseData(channel *Channel) {

}
