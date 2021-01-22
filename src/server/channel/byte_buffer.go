package channel

import (
	"fmt"
	"io"
	"math"
)

const (
	maxBufferSize = math.MaxInt32
)

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
		bufferSize:  102400,
		buffer:      make([]byte, 102400),
	}
}

func (b *ByteBuffer) Read(p []byte) (n int, err error) {
	readLength := len(p)
	readableLength := b.getReadableBytesLength()

	if readableLength == 0 {
		return 0, io.EOF
	}

	if readableLength >= readLength {
		copy(p, b.buffer[b.readerIndex:b.readerIndex+readLength])
		b.readerIndex = b.readerIndex + readLength
		return readLength, nil
	} else {
		copy(p, b.buffer[b.readerIndex:b.readerIndex+readableLength])
		b.readerIndex += readableLength
		return readableLength, nil
	}
}

func (b *ByteBuffer) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	writeLength := len(p)
	writableLength := b.getWritableBytesLength()
	if writeLength > writableLength {
		// increase the buffer size
		b.organizeSpace(b.bufferSize - b.readerIndex + 2*writeLength)
	}
	writableLength = b.getWritableBytesLength()
	// enrich the max limit
	if writeLength > writableLength {
		copy(b.buffer[b.writerIndex:b.writerIndex+writableLength], p[0:writableLength])
		b.writerIndex += writableLength
		return writableLength, nil
	} else {
		copy(b.buffer[b.writerIndex:b.writerIndex+writeLength], p[:])
		b.writerIndex += writeLength
		return writeLength, nil
	}
}

func (b *ByteBuffer) GetReaderIndex() int {
	return b.readerIndex
}

func (b *ByteBuffer) GetWriterIndex() int {
	return b.writerIndex
}

func (b *ByteBuffer) SetReaderIndex(index int) {
	b.readerIndex = index
}

func (b ByteBuffer) SetWriterIndex(index int) {
	b.writerIndex = index
}

func (b *ByteBuffer) getReadableBytesLength() int {
	return b.writerIndex - b.readerIndex
}

func (b *ByteBuffer) getWritableBytesLength() int {
	return b.bufferSize - b.writerIndex
}

func (b *ByteBuffer) organizeSpace(expect int) {
	// keep unread bytes
	readableLength := b.writerIndex - b.readerIndex
	cache := b.buffer[b.readerIndex : b.readerIndex+readableLength]
	buffersize := len(cache)
	if expect > maxBufferSize {
		buffersize = maxBufferSize
	} else {
		buffersize = expect
	}
	newBuffer := make([]byte, buffersize)
	copy(newBuffer[0:len(cache)], cache[:])
	b.buffer = newBuffer
	b.bufferSize = buffersize
	// reset the reader and writer index
	b.readerIndex = 0
	b.writerIndex = readableLength
}
