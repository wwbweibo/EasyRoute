package main

import (
	"github.com/wwbweibo/EasyRoute/example/controller"
	"github.com/wwbweibo/EasyRoute/src/http/delegate"
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"github.com/wwbweibo/EasyRoute/src/middleware"
)

func main() {
	routeContext := route.NewRouteContext()
	controller.NewHomeController(routeContext)
	routeContext.AddMiddleware(middleware.GetStaticFileMiddleware("../frontend/build", false))
	routeContext.InitRoute("/api")
	routeContext.AddDefaultHandler("/", delegate.GetDefaultDelegate("../frontend/build"))
	routeContext.AddDefaultHandler("/api", delegate.NotFoundDelegate)
	routeContext.WithServer("", "0.0.0.0", "80").Serve()
}
