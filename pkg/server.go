package pkg

import (
	"context"
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"github.com/wwbweibo/EasyRoute/pkg/route"
)

type Server struct {
	config       Config
	routeContext *route.RouteContext
	ctx          context.Context
}

func NewServer(ctx context.Context, config Config) (*Server, error) {
	routeContext := route.NewRouteContext()
	return &Server{
		config:       config,
		routeContext: routeContext,
		ctx:          ctx,
	}, nil
}

func (server *Server) AddController(controller controllers.Controller) {
	server.routeContext.AddController(controller)
}

func (server *Server) Serve() error {
	server.routeContext.InitRoute(server.config.HttpConfig.Prefix)
	return server.routeContext.Serve()
}
