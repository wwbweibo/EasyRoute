package server

import (
	"net"
)

type Channel struct {
	conn         net.Conn
	inputBuffer  *ByteBuffer
	outputBuffer *ByteBuffer
	inputChan    chan bool
	outputChan   chan bool
}

func NewChannel(conn net.Conn) *Channel {
	return &Channel{
		conn:         conn,
		inputBuffer:  NewByteBuffer(),
		outputBuffer: NewByteBuffer(),
		inputChan:    make(chan bool),
		outputChan:   make(chan bool),
	}
}

func (receiver *Channel) WriteRequestData(p []byte) {
	receiver.inputBuffer.Write(p)
	receiver.inputChan <- true
}

func (receiver *Channel) WriteResponseData(p []byte) {
	receiver.outputBuffer.Write(p)
	receiver.outputChan <- true
}
