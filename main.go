package main

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/Route"
	"reflect"
)

type ResultModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	routeContext := Route.NewRouteContext()
	NewHomeController(routeContext)
	routeContext.InitRoute(":8080")
}

type HomeController struct {
	Index func() string `Route:"/{Controller}/Index" method:"Get"`
}

func (self *HomeController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*self)
}

func NewHomeController(routeContext *Route.RouteContext) HomeController {
	instance := HomeController{
		Index: func() string {
			fmt.Println("enter index")
			return "Index"
		},
	}
	routeContext.AddController(&instance)
	return instance
}
