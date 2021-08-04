package route

import (
	controllers2 "github.com/wwbweibo/EasyRoute/controllers"
	"reflect"
	"strings"
)

// this file contains the method to resolve information from tag

// ResolveControllerName get controller name from StructTag or instance
func ResolveControllerName(controller controllers2.Controller) string {
	controllerType := controller.GetControllerType()
	controllerValue := reflect.ValueOf(controller)
	//  try to get user defined controller name, it must be controllerName field
	controllerNameValue := controllerValue.Elem().FieldByName("controllerName")
	controllerName := ""
	if controllerNameValue.IsValid() {
		controllerName = controllerNameValue.String()
	}

	// if the user defined controller name is empty, then the controller name should be ControllerTypeName
	if controllerName == "" {
		patharr := strings.Split((controllerType).String(), ".")
		controllerName = strings.Replace(patharr[len(patharr)-1], "Controller", "", 1)
	}
	return strings.ToLower(controllerName)
}

// ResolveMethodName get the method name from tag
func ResolveMethodName(field reflect.StructField) string {
	tag := field.Tag
	definedRoute := tag.Get("route")
	// the user defined route is empty, field name as default
	if definedRoute == "" {
		definedRoute = field.Name
	}
	return strings.ToLower(definedRoute)
}

// ResolveMethod get the request method from tag
func ResolveMethod(tag reflect.StructTag) string {
	method := tag.Get("method")
	// the user defined route is empty, field name as default
	if method == "" {
		method = "GET"
	}
	return method
}

// ResolveParamName get the typeName and method name map
func ResolveParamName(field reflect.StructField) []*ParamMap {
	tag := field.Tag
	paramNameString := tag.Get("param")
	if paramNameString == "" {
		return nil
	}
	paramList := make([]*ParamMap, 0)
	paramNameList := strings.Split(paramNameString, ",")
	// get method signature
	methodSignature := field.Type.String()
	methodSignature = strings.Replace(methodSignature, "func(", "", 1)
	methodSignature = methodSignature[0:strings.Index(methodSignature, ")")]
	paramType := strings.Split(methodSignature, ",")

	if len(paramType) != len(paramNameList) {
		panic("error: the method paramName and paramType not matched")
	}

	for i := 0; i < len(paramNameList); i++ {
		name := strings.Split(paramNameList[i], ":")
		m := &ParamMap{
			paramType: paramType[i],
		}

		if len(name) == 2 {
			m.paramName = name[0]
			m.source = name[1]
		} else {
			m.paramName = name[0]
			m.source = "FromQuery"
		}
		paramList = append(paramList, m)
	}

	return paramList
}
