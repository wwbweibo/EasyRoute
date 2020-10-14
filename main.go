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

	controller := NewHomeController()
	controller.RegisterAsController(routeContext)
	routeContext.RouteParse()
	routeContext.Start(":8080")
}

type HomeController struct {
	Index func() string `Route:"/{Controller}/Index" method:"Get"`
}

func (self *HomeController) RegisterAsController(ctx *Route.RouteContext) {
	ctx.AddController(self)
}

func (self *HomeController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*self)
}

func NewHomeController() HomeController {
	return HomeController{
		Index: func() string {
			fmt.Println("enter index")
			return "Index"
		},
	}
}
