package route

import "github.com/wwbweibo/EasyRoute/pkg/delegates"

type IEndPoint interface {
	MapGet(template string, handler delegates.RequestDelegate)
	MapPost(template string, handler delegates.RequestDelegate)
}

type EndPoint struct {
	Template string
	handler  delegates.RequestDelegate
	method   string
}

func (e *EndPoint) MapGet(template string, handler delegates.RequestDelegate) {

}

func (e *EndPoint) MapPost(template string, handler delegates.RequestDelegate) {

}
