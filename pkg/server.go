package pkg

import (
	"context"
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"github.com/wwbweibo/EasyRoute/pkg/route"
)

type Server struct {
	routeContext *route.RouteContext
	ctx          context.Context
}

func NewServer(ctx context.Context) (*Server, error) {
	routeContext := route.NewRouteContext()
	routeContext.InitRoute("/")

	return &Server{
		routeContext: routeContext,
		ctx:          ctx,
	}, nil
}

func (server *Server) AddController(controller controllers.Controller) {
	server.routeContext.AddController(controller)
}

func (server *Server) Serve() error {
	server.routeContext.InitRoute("/")
	return server.routeContext.Serve()
}
