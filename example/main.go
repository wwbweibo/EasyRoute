package main

import (
	"example/controller"
	"github.com/wwbweibo/EasyRoute/src/http"
	"github.com/wwbweibo/EasyRoute/src/http/route"
)

func main() {
	server := http.NewHttpServer("0.0.0.0", "80")
	routeContext := route.NewRouteContext()
	controller.NewHomeController(routeContext)
	server.RegisterHandlers(routeContext)
	server.Serve()
}
