package http

import (
	"github.com/wwbweibo/EasyRoute/server/channel"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func DecodeHttpRequest(buffer *channel.ByteBuffer) *http.Request {
	request := http.Request{}
	decodeRequestLine(buffer, &request)
	decodeRequestHeader(buffer, &request)
	decodeRequestBody(buffer, &request)
	return &request
}

// decode the http protocol request line
// return the content of request line
func decodeRequestLine(buffer *channel.ByteBuffer, request *http.Request) {
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
	requestLine = strings.Replace(requestLine, "\r\n", "", 1)
	contents := strings.Split(requestLine, " ")
	request.Method = contents[0]
	request.URL, _ = url.ParseRequestURI(contents[1])
	request.Proto = contents[2]
	protoVersions := strings.Split(strings.Replace(contents[2], "HTTP/", "", 1), ".")
	request.ProtoMajor, _ = strconv.Atoi(protoVersions[0])
	request.ProtoMinor, _ = strconv.Atoi(protoVersions[1])
}

// decode the http protocol request header
// return the map
func decodeRequestHeader(buffer *channel.ByteBuffer, request *http.Request) {
	readerIndex := buffer.GetReaderIndex()
	readedLength := 0
	request.Header = http.Header{}
	for true {
		buf := make([]byte, 1024)
		cnt, _ := buffer.Read(buf)
		content := string(buf[0:cnt])
		lines := strings.Split(content, "\r\n")
		for _, line := range lines {
			// read the empty line, end of protocol header
			if line == "" {
				buffer.SetReaderIndex(readerIndex + readedLength)
				return
			} else {
				temp := strings.Split(line, ":")
				if temp[0] == "Host" {
					request.Host = strings.Trim(temp[1], " ")
				} else {
					request.Header.Add(temp[0], strings.Trim(temp[1], " "))
				}
				readedLength += len(line) + 2
			}
		}
	}
}

func decodeRequestBody(buffer *channel.ByteBuffer, request *http.Request) {
	// todo decode request body
	if request.Method == http.MethodGet {
		request.Body = nil
	} else {
		// todo will read body
		var buf = make([]byte, 0)
		for true {
			bytes := make([]byte, 1024)
			cnt, _ := buffer.Read(bytes)
			buf = append(buf, bytes[0:cnt]...)
			if cnt < 1024 {
				return
			}
		}
	}
}
