package main

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/route"
	"reflect"
)

type ResultModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	routeContext := route.NewRouteContext()
	routeContext.RegisterTypeByInstance(Person{})
	outerMiddleware := func(next route.RequestDelegate) route.RequestDelegate {
		return func(ctx route.HttpContext) {
			fmt.Println("abc")
			next(ctx)
		}
	}

	routeContext.AddMiddleware(outerMiddleware)

	routeContext.AddMiddleware(
		func(next route.RequestDelegate) route.RequestDelegate {
			return func(ctx route.HttpContext) {
				fmt.Println("before")
				next(ctx)
				fmt.Println("after")
			}
		},
	)

	NewHomeController(routeContext)
	routeContext.InitRoute(":8080")
}

type Person struct {
	Name string  `json:"Name"`
	Age  float64 `json:"Age"`
}

type HomeController struct {
	Index       func() string              `method:"Get"`
	IndexA      func(a string) string      `method:"Get" param:"a"`
	IndexPerson func(person Person) Person `method:"get" param:"person"`
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
		IndexA: func(a string) string {
			fmt.Println(a)
			return a
		},
		IndexPerson: func(person Person) Person {
			fmt.Printf("Name： %s, Age：%f\n", person.Name, person.Age)
			return person
		},
	}
	routeContext.AddController(&instance)
	return instance
}
