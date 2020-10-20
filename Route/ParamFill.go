package Route

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

// fill up the param list
func fillUp(c *gin.Context, paramList *[]paramMap) []reflect.Value {
	paramValueList := make([]reflect.Value, len(*paramList))
	for idx, param := range *paramList {
		if param.source == "FromQuery" {
			paramValueList[idx] = fillUpFromQueryString(c, &param)
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
		v, _ := typeCollectionInstance.InstanceOf(param.paramType)
		vv := reflect.New(*paramType).Elem()
		err := json.Unmarshal([]byte(value), &v)
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < vv.NumField(); i++ {
			field := vv.Field(i)
			fieldValue := v.(map[string]interface{})[(*paramType).Field(i).Name]
			field.Set(reflect.ValueOf(fieldValue))
		}

		return vv
	}
	return reflect.ValueOf(value)
}
