package route

import (
	"reflect"
	"strings"
)

// will scan all registered controllers, and Convert each method to a request delegate
func scanEndPoint(routeContext *RouteContext, prefix string) {
	for _, controller := range routeContext.controllers {
		controllerType := (*controller).GetControllerType()
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
