package route

type IEndPoint interface {
	MapGet(template string, handler RequestDelegate)
	MapPost(template string, handler RequestDelegate)
}

type EndPoint struct {
	Template string
	handler  RequestDelegate
	method   string
}

func (e *EndPoint) MapGet(template string, handler RequestDelegate) {

}

func (e *EndPoint) MapPost(template string, handler RequestDelegate) {

}
