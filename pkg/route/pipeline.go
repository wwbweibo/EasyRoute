package route

import (
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
)

// http 处理管道

type Pipeline struct {
	handlerList []delegates.Middleware
}

func (receiver *Pipeline) AddMiddleware(middleware delegates.Middleware) {
	receiver.handlerList = append(receiver.handlerList, middleware)
}

func (receiver *Pipeline) build() delegates.RequestDelegate {
	var app delegates.RequestDelegate
	app = reqHandler.delegate
	for i := len(receiver.handlerList) - 1; i >= 0; i-- {
		app = receiver.handlerList[i](app)
	}
	return app
}
