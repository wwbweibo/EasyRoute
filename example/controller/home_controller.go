package controller

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/pkg/route"
	"reflect"
)

type HomeController struct {
	Index func() string `method:"Get"`
}

func (self *HomeController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*self)
}

func NewHomeController(routeContext *route.RouteContext) HomeController {
	instance := HomeController{
		Index: func() string {
			fmt.Println("enter index")
			return "Index"
		},
	}
	routeContext.AddController(&instance)
	return instance
}
