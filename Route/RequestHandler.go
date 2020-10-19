package Route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type requestHandler struct {
	routeContext *RouteContext
	delegate     RequestDelegate
}

var reqHandler = requestHandler{
	routeContext: &routeContext,
	delegate: func(ctx HttpContext) {
		c := (*gin.Context)(ctx)
		path := c.Request.URL.Path
		if routeMap, ok := routeContext.routeMap[path]; ok {
			if c.Request.Method == routeMap.method {
				methodName := strings.Replace(path, "/"+routeMap.controllerName+"/", "", 1)
				method := reflect.ValueOf(*routeMap.controller).Elem().FieldByName(methodName)
				// if the length of param map greater than 0, the method got params, fill it
				if len(*routeMap.paramMap) > 0 {
					result := method.Call(fillUp(c, routeMap.paramMap))[0]
					c.String(http.StatusOK, result.String())
				} else {
					result := method.Call(nil)[0]
					c.String(http.StatusOK, result.String())
				}
			} else {
				c.String(http.StatusMethodNotAllowed, "405 NotAllowed")
			}

		} else {
			c.String(http.StatusNotFound, "404 NotFind")
		}
	},
}
