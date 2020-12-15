package http

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/server/channel"
	"strings"
)

func DecodeHttpRequest(buffer *channel.ByteBuffer) {
	requestLine := decodeRequestLine(buffer)
	fmt.Println(requestLine)
	decodeRequestHeader(buffer)
	decodeRequestBody(buffer)
}

// decode the http protocol request line
// return the content of request line
func decodeRequestLine(buffer *channel.ByteBuffer) string {
	readerIndex := buffer.GetReaderIndex()
	flag := true
	var requestLine string
	for flag {
		buf := make([]byte, 1024)
		cnt, _ := buffer.Read(buf)
		content := string(buf[0:cnt])
		// the request line end with '\r\n'
		if strings.Contains(content, "\r\n") {
			idx := strings.Index(content, "\r\n")
			requestLine += content[0 : idx+2]
			buffer.SetReaderIndex(readerIndex + len(requestLine))
			flag = false
		} else {
			requestLine += content
		}
	}
	return requestLine
}

// decode the http protocol request header
// return the map
func decodeRequestHeader(buffer *channel.ByteBuffer) map[string]string {
	readerIndex := buffer.GetReaderIndex()
	readedLength := 0
	requestHeader := make(map[string]string)
	for true {
		buf := make([]byte, 1024)
		cnt, _ := buffer.Read(buf)
		content := string(buf[0:cnt])
		lines := strings.Split(content, "\r\n")
		for _, line := range lines {
			// read the empty line, end of protocol header
			if line == "" {
				buffer.SetReaderIndex(readerIndex + readedLength)
				for k, v := range requestHeader {
					fmt.Println(k + ":" + v)
				}
				return requestHeader
			} else {
				temp := strings.Split(line, ":")
				requestHeader[temp[0]] = temp[1]
				readedLength += len(line) + 2
			}
		}
	}
	return nil
}

func decodeRequestBody(buffer *channel.ByteBuffer) []byte {
	buf := make([]byte, 1024)
	cnt, _ := buffer.Read(buf)
	content := string(buf[0:cnt])
	fmt.Println(content)
	return buf[0:cnt]
}
