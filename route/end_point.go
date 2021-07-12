package route

import (
	delegates2 "github.com/wwbweibo/EasyRoute/delegates"
)

type IEndPoint interface {
	MapGet(template string, handler delegates2.RequestDelegate)
	MapPost(template string, handler delegates2.RequestDelegate)
}

type EndPoint struct {
	Template string
	handler  delegates2.RequestDelegate
	method   string
}

func (e *EndPoint) MapGet(template string, handler delegates2.RequestDelegate) {

}

func (e *EndPoint) MapPost(template string, handler delegates2.RequestDelegate) {

}
