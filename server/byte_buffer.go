package server

type ByteBuffer struct {
	readerIndex int
	writerIndex int
	bufferSize  int
	buffer      []byte
}

func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		readerIndex: 0,
		writerIndex: 0,
		buffer:      make([]byte, 0),
	}
}

func (b *ByteBuffer) Read(p []byte) (n int, err error) {
	if len(b.buffer) < len(p) {
		length := len(b.buffer)
		copy(p, b.buffer)
		b.buffer = []byte{}
		return length, nil
	} else {
		copy(p, b.buffer[0:len(p)])
		return len(p), nil
	}
}

func (b *ByteBuffer) Write(p []byte) (n int, err error) {
	b.buffer = append(b.buffer, p...)
	return len(p), nil
}
