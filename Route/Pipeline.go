package Route

type Pipeline struct {
	handlerList []Middleware
}

func (receiver *Pipeline) AddMiddleware(middleware Middleware) {
	receiver.handlerList = append(receiver.handlerList, middleware)
}

func (receiver *Pipeline) build() RequestDelegate {
	var app RequestDelegate
	app = reqHandler.delegate
	for i := len(receiver.handlerList) - 1; i >= 0; i-- {
		app = receiver.handlerList[i](app)
	}
	return app
}

type RequestDelegate func(ctx HttpContext)

type Middleware func(next RequestDelegate) RequestDelegate
