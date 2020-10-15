package Route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type requestHandler struct {
	routeContext *RouteContext
}

func (self *requestHandler) Handle(ctx RequestContext, next PipelineHandler) {
	c := (*gin.Context)(ctx)
	path := c.Request.RequestURI
	if routeMap, ok := self.routeContext.routeMap[path]; ok {
		if c.Request.Method == routeMap.method {
			methodName := strings.Replace(path, "/"+routeMap.controllerName+"/", "", 1)
			result := reflect.ValueOf(*routeMap.controller).Elem().FieldByName(methodName).Call(nil)[0]
			c.String(http.StatusOK, result.String())
		} else {
			c.String(http.StatusMethodNotAllowed, "405 NotAllowed")
		}

	} else {
		c.String(http.StatusNotFound, "404 NotFind")
	}
}
