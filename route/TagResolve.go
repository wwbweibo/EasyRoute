package route

import (
	"reflect"
	"strings"
)

// this file contains the method to resolve information from tag

// get controller name from StructTag or instance
func resolveControllerName(controllerType *reflect.Type, controller *Controller) string {
	//  try to get user defined controller name, it must be controllerName field
	controllerNameValue := reflect.ValueOf(controller).Elem().Elem().Elem().FieldByName("controllerName")
	controllerName := ""
	if controllerNameValue.IsValid() {
		controllerName = controllerNameValue.String()
	}

	// if the user defined controller name is empty, then the controller name should be ControllerTypeName
	if controllerName == "" {
		patharr := strings.Split((*controllerType).String(), ".")
		controllerName = strings.Replace(patharr[len(patharr)-1], "Controller", "", 1)
	}
	return controllerName
}

// get the method name from tag
func resolveMethodName(tag *reflect.StructTag, field *reflect.StructField) string {
	definedRoute := (*tag).Get("route")
	// the user defined route is empty, field name as default
	if definedRoute == "" {
		definedRoute = field.Name
	}
	return definedRoute
}

// get the request method from tag
func resolveMethod(tag *reflect.StructTag) string {
	method := (*tag).Get("Method")
	// the user defined route is empty, field name as default
	if method == "" {
		method = "GET"
	}
	return method
}

// get the typeName and method name map
func resolveParamName(tag *reflect.StructTag, field *reflect.StructField) *[]paramMap {
	paramNameString := tag.Get("param")
	if paramNameString == "" {
		return nil
	}
	paramList := make([]paramMap, 0)
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
		m := paramMap{
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

	return &paramList
}
