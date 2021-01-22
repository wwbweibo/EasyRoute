package main

import (
	"github.com/wwbweibo/EasyRoute/example/controller"
	"github.com/wwbweibo/EasyRoute/src/http"
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"github.com/wwbweibo/EasyRoute/src/middleware"
)

func main() {
	server := http.NewHttpServer("0.0.0.0", "80")
	routeContext := route.NewRouteContext()
	routeContext.AddMiddleware(middleware.GetStaticFileMiddleware("../frontend/build", false))

	controller.NewHomeController(routeContext)
	server.RegisterHandlers(routeContext)
	server.Serve()
}
