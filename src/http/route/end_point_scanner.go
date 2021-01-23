package route

import (
	"reflect"
	"strings"
)

// will scan all registered controllers, and Convert each method to a request delegate
func scanEndPoint(routeContext *RouteContext, prefix string) {
	for _, controller := range routeContext.controllers {
		controllerType := (*controller).GetControllerType()
		controllerName := resolveControllerName(&controllerType, controller)
		for i := 0; i < controllerType.NumField(); i++ {
			field := controllerType.Field(i)
			route := resolveMethodName(&field.Tag, &field)
			method := resolveMethod(&field.Tag)
			paramList := resolveParamName(&field.Tag, &field)

			// if the route is not start with "/", then combine the controllerName and route
			if !strings.HasPrefix(route, "/") {
				route = "/" + controllerName + "/" + route
			}
			if prefix != "" {
				route = prefix + route
			}

			methodValue := reflect.ValueOf(*controller).Elem().Field(i)

			requestHandler := convertControllerMethodToRequestDelegate(methodValue, paramList, method)

			endpoint := &EndPoint{
				Template: route,
				method:   method,
				handler:  requestHandler,
			}

			routeContext.endPointTrie.AddEndPoint(endpoint)
		}
	}
}
