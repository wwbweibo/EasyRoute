package channel

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

func (receiver *Channel) WriteResponseData(p []byte) (int, error) {
	n, err := receiver.outputBuffer.Write(p)
	receiver.outputChan <- true
	return n, err
}

func (receiver *Channel) GetInputBuffer() *ByteBuffer {
	return receiver.inputBuffer
}

func (receiver *Channel) GetOutputBuffer() *ByteBuffer {
	return receiver.outputBuffer
}

func (receiver *Channel) GetInputChannel() chan bool {
	return receiver.inputChan
}

func (receiver *Channel) GetOutputChannel() chan bool {
	return receiver.outputChan
}

func (receiver *Channel) GetConnection() net.Conn {
	return receiver.conn
}

func (receiver *Channel) Write(p []byte) (n int, err error) {
	return receiver.WriteResponseData(p)
}
