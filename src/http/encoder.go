package http

import (
	"github.com/wwbweibo/EasyRoute/src/server/channel"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func DecodeHttpRequest(buffer *channel.ByteBuffer) *http.Request {
	request := http.Request{}
	readerindex := buffer.GetReaderIndex()
	if !decodeRequestLine(buffer, &request) {
		buffer.SetReaderIndex(readerindex)
		return nil
	}
	if !decodeRequestHeader(buffer, &request) {
		buffer.SetReaderIndex(readerindex)
		return nil
	}
	if !decodeRequestBody(buffer, &request) {
		buffer.SetReaderIndex(readerindex)
		return nil
	}
	return &request
}

// decode the http protocol request line
// return the content of request line
func decodeRequestLine(buffer *channel.ByteBuffer, request *http.Request) bool {
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
	return true
}

// decode the http protocol request header
// return the map
func decodeRequestHeader(buffer *channel.ByteBuffer, request *http.Request) bool {
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
				buffer.SetReaderIndex(readerIndex + readedLength + 2)
				return true
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
	return true
}

func decodeRequestBody(buffer *channel.ByteBuffer, request *http.Request) bool {
	// todo decode request body
	if request.Method == http.MethodGet {
		request.Body = nil
		return true
	} else {
		contentType := request.Header["Content-Type"][0]
		if strings.Contains(contentType, "multipart/form-data;") {
			return decodeMultipartFormData(buffer, request)
		}
	}
	return false
}

func decodeMultipartFormData(buffer *channel.ByteBuffer, request *http.Request) bool {
	// boundary=--------------------------914951608272316434835454
	contentType := request.Header["Content-Type"][0]
	contentLength, _ := strconv.Atoi(request.Header["Content-Length"][0])
	var boundary string
	for _, c := range strings.Split(strings.Replace(contentType, "multipart/form-data;", "", 1), ";") {
		cc := strings.Split(c, "=")
		if strings.Trim(cc[0], " ") == "boundary" {
			boundary = cc[1]
		}
	}
	readableBytes := buffer.GetWriterIndex() - buffer.GetReaderIndex()
	if readableBytes < contentLength {
		return false
	}
	reader := multipart.NewReader(buffer, boundary)
	form, _ := reader.ReadForm(int64(contentLength))

	request.MultipartForm = form
	return true
}

//----------------------------914951608272316434835454
//Content-Disposition: form-data; name="aa"
//
//a
