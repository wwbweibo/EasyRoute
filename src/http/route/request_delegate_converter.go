package route

import (
	"encoding/json"
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"net/http"
	"reflect"
)

func convertControllerMethodToRequestDelegate(method reflect.Value, params []*paramMap, httpMethod string) RequestDelegate {
	delegate := func(ctx *context.Context) {
		request := ctx.Request
		if request.Method == httpMethod {
			// if the length of param map greater than 0, the method got params, fill it
			if params == nil || len(params) <= 0 {
				result := method.Call(nil)[0]
				data, _ := json.Marshal(result.Interface())
				ctx.Response.WriteBody(data)
				ctx.Response.WriteHttpCode(http.StatusOK, "OK")
			} else {
				params := fillUp(request, params)
				result := method.Call(params)[0]
				data, _ := json.Marshal(result.Interface())
				ctx.Response.WriteBody(data)
				ctx.Response.WriteHttpCode(http.StatusOK, "OK")
			}
		} else {
			ctx.Response.WriteHttpCode(http.StatusMethodNotAllowed, "MethodNotAllowed")
		}
	}
	return delegate
}
