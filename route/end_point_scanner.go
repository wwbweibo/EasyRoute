package route

import (
	controllers2 "github.com/wwbweibo/EasyRoute/controllers"
	"github.com/wwbweibo/EasyRoute/logger"
	"reflect"
	"strings"
)

// ScanEndPoint will scan all registered controllers, and Convert each method to a request delegate
func ScanEndPoint(endPointTrie *EndPointTrie, controllers []controllers2.Controller, prefix string) {
	for _, controller := range controllers {
		controllerType := controller.GetControllerType()
		logger.Info("[route] - [ScanEndPoint] scan end point on controller %s", controllerType.String())
		controllerName := ResolveControllerName(controller)
		for i := 0; i < controllerType.NumField(); i++ {
			field := controllerType.Field(i)
			route := ResolveMethodName(field)
			method := ResolveMethod(field.Tag)
			paramList := ResolveParamName(field)

			// if the route is not start with "/", then combine the controllerName and route
			if !strings.HasPrefix(route, "/") {
				route = "/" + controllerName + "/" + route
			}
			if prefix != "" && prefix != "/" {
				route = prefix + route
			}
			logger.Info("[route] - [ScanEndPoint] find route %s", route)

			// get the method body and convert it to a request delegate
			methodValue := reflect.ValueOf(controller).Elem().Field(i)
			requestHandler := convertControllerMethodToRequestDelegate(methodValue, paramList, method)

			// init end point
			endpoint := &EndPoint{
				Template: route,
				method:   method,
				handler:  requestHandler,
			}

			// and it to tree
			endPointTrie.AddEndPoint(endpoint)
		}
	}
}
