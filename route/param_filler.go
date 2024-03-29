//package route
//
//import (
//	"bytes"
//	"encoding/json"
//	http3 "github.com/wwbweibo/EasyRoute/http"
//	"github.com/wwbweibo/EasyRoute/logger"
//	"net/http"
//	"reflect"
//	"strings"
//)
//
//type ParamMap struct {
//	paramName string
//	paramType string
//	source    string
//}
//
//// 参数填充方法
//
//// fill up the param list
//func fillUp(ctx *http3.Context, paramList []ParamMap) []reflect.Value {
//	paramValueList := make([]reflect.Value, len(paramList))
//	if paramList[0].paramType == "context.Context" {
//		paramValueList[0] = reflect.ValueOf(ctx.Ctx)
//	}
//	request := ctx.Request
//	for idx, param := range paramList {
//		if param.source == "FromQuery" {
//			paramValueList[idx] = fillUpFromQueryString(request, param)
//		} else if param.source == "FromBody" {
//			paramValueList[idx] = fillUpFromBodyContent(request, param)
//		} else if param.source == "FromForm" {
//			paramValueList[idx] = fillUpFromForm(request, param)
//		}
//	}
//	return paramValueList
//}
//
//// fill up params from query string
//func fillUpFromQueryString(request *http.Request, param ParamMap) reflect.Value {
//	value := request.URL.Query().Get(param.paramName)
//	paramType, err := typeCollectionInstance.TypeOf(param.paramType)
//	if err != nil {
//		panic("Error to Fill Up Param " + param.paramName + " could not find the type")
//	}
//
//	// if type of struct, need to deserialize
//	if (paramType).Kind() == reflect.Struct {
//		return deserializeJsonObject(param.paramType, []byte(value))
//	}
//	return reflect.ValueOf(value)
//}
//
//func fillUpFromBodyContent(request *http.Request, param ParamMap) reflect.Value {
//	bodyContent := request.Body
//	buf := bytes.Buffer{}
//	buf.ReadFrom(bodyContent)
//	return deserializeJsonObject(param.paramType, buf.Bytes())
//}
//
//func fillUpFromForm(request *http.Request, param ParamMap) reflect.Value {
//	//
//	//err := request.ParseForm()
//	//if err != nil {
//	//	panic("error to parse form " + err.Error())
//	//}
//
//	value := request.PostFormValue(param.paramName) // request.Form[param.paramName]
//	if len(value) <= 0 {
//		return reflect.ValueOf(nil)
//	}
//	return deserializeJsonObject(param.paramType, []byte(value))
//}
//
//func fillUpFromPath(request http.Request, param ParamMap) reflect.Value {
//	// todo : path 变量的获取方式
//	url := request.URL.String()
//	pathList := strings.Split(url, "/")
//	value := pathList[len(pathList)-1]
//	return reflect.ValueOf(value)
//}
//
//func deserializeJsonObject(typeName string, jsonData []byte) reflect.Value {
//	instance, err := typeCollectionInstance.InstanceOf(typeName)
//	if err != nil {
//		panic("error to handle request")
//	}
//	err = json.Unmarshal(jsonData, instance.Interface())
//	if err != nil {
//		logger.Error("error to unmarshal data", err)
//		panic("error to handle request, data error")
//	}
//	return instance.Elem()
//}

package route
