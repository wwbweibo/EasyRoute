package route

import (
	"encoding/json"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	http2 "github.com/wwbweibo/EasyRoute/pkg/http"
	"net/http"
	"reflect"
)

func convertControllerMethodToRequestDelegate(method reflect.Value, params []*paramMap, httpMethod string) delegates.RequestDelegate {
	delegate := func(ctx *http2.HttpContext) {
		request := ctx.Request
		if request.Method == httpMethod {
			// if the length of param map greater than 0, the method got params, fill it
			if params == nil || len(params) <= 0 {
				result := method.Call(nil)[0]
				data, _ := json.Marshal(result.Interface())
				ctx.Response.Write(data)
				ctx.Response.WriteHeader(http.StatusOK)
			} else {
				params := fillUp(request, params)
				result := method.Call(params)[0]
				data, _ := json.Marshal(result.Interface())
				ctx.Response.Write(data)
				ctx.Response.WriteHeader(http.StatusOK)
			}
		} else {
			ctx.Response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
	return delegate
}
