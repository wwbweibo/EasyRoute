package route

import (
	delegates2 "github.com/wwbweibo/EasyRoute/delegates"
	"github.com/wwbweibo/EasyRoute/log"
)

// http 处理管道

type Pipeline struct {
	handlerList []delegates2.Middleware
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		handlerList: []delegates2.Middleware{},
	}
}

func (receiver *Pipeline) AddMiddleware(middleware delegates2.Middleware) {
	receiver.handlerList = append(receiver.handlerList, middleware)
}

func (receiver *Pipeline) Build(trie *EndPointTrie) delegates2.RequestDelegate {
	log.Info("[route] - [Pipeline] - [Build] building request pipeline", nil)
	var app delegates2.RequestDelegate
	app = newRequestHandler(trie).delegate
	for i := len(receiver.handlerList) - 1; i >= 0; i-- {
		app = receiver.handlerList[i](app)
	}
	return app
}
