package route

import (
	"encoding/json"
	"github.com/wwbweibo/EasyRoute/http/context"
	"net/http"
	"reflect"
	"strings"
)

// http请发分发

type requestHandler struct {
	routeContext *RouteContext
	delegate     RequestDelegate
}

var reqHandler = requestHandler{
	routeContext: &routeContext,
	delegate: func(ctx *context.Context) {
		request := ctx.Request
		path := request.URL.Path
		if routeMap, ok := routeContext.routeMap[path]; ok {
			if request.Method == routeMap.method {
				methodName := strings.Replace(path, "/"+routeMap.controllerName+"/", "", 1)
				method := reflect.ValueOf(*routeMap.controller).Elem().FieldByName(methodName)
				// if the length of param map greater than 0, the method got params, fill it
				if routeMap.paramMap == nil || len(*routeMap.paramMap) <= 0 {
					result := method.Call(nil)[0]
					data, _ := json.Marshal(result.Interface())
					ctx.Response.WriteBody(data)
					ctx.Response.WriteHttpCode(200)
				} else {
					params := fillUp(request, routeMap.paramMap)
					result := method.Call(params)[0]
					data, _ := json.Marshal(result.Interface())
					ctx.Response.WriteBody(data)
					ctx.Response.WriteHttpCode(200)
				}
			} else {
				ctx.Response.WriteHttpCode(http.StatusMethodNotAllowed)
			}

		} else {
			ctx.Response.WriteHttpCode(http.StatusNotFound)
		}
	},
}
