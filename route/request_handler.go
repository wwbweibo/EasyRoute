package route

type requestHandler struct {
	routeContext *RouteContext
	delegate     RequestDelegate
}

var reqHandler = requestHandler{
	//routeContext: &routeContext,
	//delegate: func(ctx http.Context) {
	//	request := ctx.Request
	//	path := request.URL.Path
	//	if routeMap, ok := routeContext.routeMap[path]; ok {
	//		if request.Method == routeMap.method {
	//			methodName := strings.Replace(path, "/"+routeMap.controllerName+"/", "", 1)
	//			method := reflect.ValueOf(*routeMap.controller).Elem().FieldByName(methodName)
	//			// if the length of param map greater than 0, the method got params, fill it
	//			if len(*routeMap.paramMap) > 0 {
	//				//params := fillUp(request, routeMap.paramMap)
	//				//result := method.Call(params)[0]
	//				//request.Response.
	//				//c.JSON(http2.StatusOK, result.Interface())
	//			} else {
	//				//result := method.Call(nil)[0]
	//				//c.JSON(http2.StatusOK, result)
	//			}
	//		} else {
	//			// c.String(http.StatusMethodNotAllowed, "405 NotAllowed")
	//		}
	//
	//	} else {
	//
	//	}
	//},
}
