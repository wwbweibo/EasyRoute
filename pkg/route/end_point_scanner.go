package route

import (
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"reflect"
	"strings"
)

// ScanEndPoint will scan all registered controllers, and Convert each method to a request delegate
func ScanEndPoint(endPointTrie *EndPointTrie, controllers []controllers.Controller, prefix string) {
	for _, controller := range controllers {
		controllerType := controller.GetControllerType()
		controllerName := ResolveControllerName(&controllerType, controller)
		for i := 0; i < controllerType.NumField(); i++ {
			field := controllerType.Field(i)
			route := ResolveMethodName(&field.Tag, &field)
			method := ResolveMethod(&field.Tag)
			paramList := ResolveParamName(&field.Tag, &field)

			// if the route is not start with "/", then combine the controllerName and route
			if !strings.HasPrefix(route, "/") {
				route = "/" + controllerName + "/" + route
			}
			if prefix != "" && prefix != "/" {
				route = prefix + route
			}

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
