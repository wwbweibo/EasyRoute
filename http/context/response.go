package context

import (
	"io"
	"net/http"
)

type Response struct {
	resp *http.Response
}

type responseBody struct {
	data []byte
}

func (r responseBody) Read(p []byte) (n int, err error) {
	length := len(r.data)
	if len(p) < length {
		copy(p, r.data[0:len(p)])
		r.data = r.data[len(p)+1:]
		return len(p), nil
	} else {
		copy(p, r.data)
		return length, nil
	}
}

func (r responseBody) Close() error {
	r.data = nil
	return nil
}

func NewResponse() *Response {
	return &Response{resp: &http.Response{}}
}

func (response Response) GetHeader() http.Header {
	return response.resp.Header
}

func (response Response) WriteHeader(key string, value []string) {
	response.resp.Header[key] = value
}

func (response Response) WriteHttpCode(code int) {
	response.resp.StatusCode = code

}

func (response Response) WriteBody(p []byte) (int, error) {
	body := responseBody{
		data: p,
	}
	response.resp.Body = body
	response.resp.ContentLength = int64(len(p))
	return len(p), nil
}

func (response Response) Write(w io.Writer) {
	response.resp.Write(w)
}
