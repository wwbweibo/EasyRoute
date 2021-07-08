package route

import (
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
)

var routeContext = RouteContext{
	controllers: make([]*controllers.Controller, 0),

	pipeline: Pipeline{
		handlerList: make([]delegates.Middleware, 0),
	},
	typeCollection: typeCollectionInstance,
	endPointTrie:   NewEndPointTrie(),
}
