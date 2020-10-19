package Route

import (
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

func fillUpFromQueryString(c *gin.Context, param *paramMap) reflect.Value {
	value := c.Query(param.paramName)
	return reflect.ValueOf(value)
}
