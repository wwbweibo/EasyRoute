package Route

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// fill up the param list
func fillUp(c *gin.Context, paramList *[]paramMap) []reflect.Value {
	paramValueList := make([]reflect.Value, len(*paramList))
	for idx, param := range *paramList {
		if param.source == "FromQuery" {
			paramValueList[idx] = fillUpFromQueryString(c, &param)
		} else if param.source == "FromBody" {
			paramValueList[idx] = fillUpFromBodyContent(c, &param)
		} else if param.source == "FromForm" {
			paramValueList[idx] = fillUpFromForm(c, &param)
		}
	}
	return paramValueList
}

// fill up params from query string
func fillUpFromQueryString(c *gin.Context, param *paramMap) reflect.Value {
	value := c.Query(param.paramName)
	paramType, err := typeCollectionInstance.TypeOf(param.paramType)
	if err != nil {
		panic("Error to Fill Up Param " + param.paramName + " could not find the type")
	}

	// if type of struct, need to deserialize
	if (*paramType).Kind() == reflect.Struct {
		return deserializeJsonObject(param.paramType, []byte(value))
	}
	return reflect.ValueOf(value)
}

func fillUpFromBodyContent(c *gin.Context, param *paramMap) reflect.Value {
	bodyContent := c.Request.Body
	buf := bytes.Buffer{}
	buf.ReadFrom(bodyContent)
	return deserializeJsonObject(param.paramType, buf.Bytes())
}

func fillUpFromForm(c *gin.Context, param *paramMap) reflect.Value {
	value := c.Request.Form[param.paramName]
	if len(value) <= 0 {
		return reflect.ValueOf(nil)
	}
	return deserializeJsonObject(param.paramType, []byte(value[0]))
}

func fillUpFromPath(c *gin.Context, param *paramMap) reflect.Value {
	// todo : path 变量的获取方式
	url := c.Request.URL.String()
	pathList := strings.Split(url, "/")
	value := pathList[len(pathList)-1]
	return reflect.ValueOf(value)
}

func deserializeJsonObject(typeName string, jsonData []byte) reflect.Value {
	t, _ := typeCollectionInstance.TypeOf(typeName)
	instance, err := typeCollectionInstance.InstanceOf(typeName)
	if err != nil {
		panic("error to handle request")
	}
	instance2 := reflect.New(*t)
	err = json.Unmarshal(jsonData, &instance)
	if err != nil {
		panic("error to handle request, data error")
	}
	for i := 0; i < instance2.NumField(); i++ {
		field := instance2.Field(i)
		fieldValue := instance.(map[string]interface{})[(*t).Field(i).Name]
		field.Set(reflect.ValueOf(fieldValue))
	}
	return instance2
}
