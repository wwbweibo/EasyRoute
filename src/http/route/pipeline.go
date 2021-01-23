package route

import (
	"github.com/wwbweibo/EasyRoute/src/http"
)

// http 处理管道

type Pipeline struct {
	handlerList []Middleware
}

func (receiver *Pipeline) AddMiddleware(middleware Middleware) {
	receiver.handlerList = append(receiver.handlerList, middleware)
}

func (receiver *Pipeline) build() http.RequestDelegate {
	var app http.RequestDelegate
	app = reqHandler.delegate
	for i := len(receiver.handlerList) - 1; i >= 0; i-- {
		app = receiver.handlerList[i](app)
	}
	return app
}

type Middleware func(next http.RequestDelegate) http.RequestDelegate
