package main

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/example/controller"
	"github.com/wwbweibo/EasyRoute/src/http/delegate"
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"github.com/wwbweibo/EasyRoute/src/middleware"
)

func main() {
	routeContext := route.NewRouteContext()
	controller.NewHomeController(routeContext)
	controller.NewUserController(routeContext)
	routeContext.RegisterTypeByInstance(controller.User{
		Username: "",
		Password: "",
	})
	routeContext.AddMiddleware(middleware.GetStaticFileMiddleware("/home/weibo/data/Code/Go/EasyRoute/frontend/build", false))
	routeContext.InitRoute("/api")
	routeContext.AddDefaultHandler("/", delegate.GetDefaultDelegate("/home/weibo/data/Code/Go/EasyRoute/frontend/build"))
	routeContext.AddDefaultHandler("/api", delegate.NotFoundDelegate)
	err := routeContext.WithServer("", "0.0.0.0", "80").Serve()
	if err != nil {
		fmt.Println(err.Error())
	}
}
