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
	outerMiddleware := func(next Route.RequestDelegate) Route.RequestDelegate {
		return func(ctx Route.HttpContext) {
			fmt.Println("abc")
			next(ctx)
		}
	}

	routeContext.AddMiddleware(outerMiddleware)

	routeContext.AddMiddleware(
		func(next Route.RequestDelegate) Route.RequestDelegate {
			return func(ctx Route.HttpContext) {
				fmt.Println("before")
				next(ctx)
				fmt.Println("after")
			}
		},
	)

	NewHomeController(routeContext)
	routeContext.InitRoute(":8080")
}

type HomeController struct {
	Index  func() string         `method:"Get"`
	IndexA func(a string) string `method:"Get" param:"a"`
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
		IndexA: func(a string) string {
			fmt.Println(a)
			return a
		},
	}
	routeContext.AddController(&instance)
	return instance
}
