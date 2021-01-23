package route

import "github.com/wwbweibo/EasyRoute/src/http"

type IEndPoint interface {
	MapGet(template string, handler http.RequestDelegate)
	MapPost(template string, handler http.RequestDelegate)
}

type EndPoint struct {
	Template string
	handler  http.RequestDelegate
	method   string
}

func (e *EndPoint) MapGet(template string, handler http.RequestDelegate) {

}

func (e *EndPoint) MapPost(template string, handler http.RequestDelegate) {

}
