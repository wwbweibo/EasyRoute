package route

import (
	delegates2 "github.com/wwbweibo/EasyRoute/delegates"
	http3 "github.com/wwbweibo/EasyRoute/http"
	"net/http"
	"reflect"
	"strings"
)

func convertControllerMethodToRequestDelegate(method reflect.Value, params []ParamMap, httpMethod string) delegates2.RequestDelegate {
	delegate := func(ctx *http3.Context) {
		request := ctx.Request
		if strings.ToLower(request.Method) == strings.ToLower(httpMethod) {
			// if the length of param map greater than 0, the method got params, fill it
			if params == nil || len(params) <= 0 {
				result := method.Call(nil)[0]
				_ = ctx.WriteJson(result.Interface(), http.StatusOK)
			} else {
				params := fillUp(request, params)
				result := method.Call(params)[0]
				_ = ctx.WriteJson(result.Interface(), http.StatusOK)
				// ctx.Response.Write(data)
			}
		} else {
			ctx.Response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
	return delegate
}
