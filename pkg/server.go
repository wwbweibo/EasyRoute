package pkg

import (
	"context"
	"errors"
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	"github.com/wwbweibo/EasyRoute/pkg/route"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
)

type Server struct {
	config       Config
	ctx          context.Context
	controllers  []controllers.Controller  // controllers is all the Controller add to current server
	pipeline     *route.Pipeline           // pipeline defines the request pipeline for current server
	app          delegates.RequestDelegate // app is the entry point for a request
	endPointTrie *route.EndPointTrie       // endPointTrie maintain all resolved endpoint
}

func NewServer(ctx context.Context, config Config) (*Server, error) {
	// routeContext := route.NewRouteContext(ctx)
	return &Server{
		config:       config,
		ctx:          ctx,
		controllers:  []controllers.Controller{},
		pipeline:     route.NewPipeline(),
		endPointTrie: route.NewEndPointTrie(),
	}, nil
}

// AddController will register the given handle to server
func (server *Server) AddController(controller controllers.Controller) {
	server.controllers = append(server.controllers, controller)
}

// AddMiddleware will add a middle to pipeline
func (server *Server) AddMiddleware(middleware delegates.Middleware) {
	server.pipeline.AddMiddleware(middleware)
}

// AddDefaultHandler add a default action for given pattern
func (server *Server) AddDefaultHandler(pattern string, delegate delegates.RequestDelegate) {
	targetNode, _ := server.endPointTrie.GetRoot().Search(strings.Split(pattern, "/")[1:])
	targetNode.DefaultHandler = delegate
}

// Serve will make server ready for request and listen with given config
func (server *Server) Serve() error {
	server.prepare()
	handler := &route.HttpRequestHandler{
		RequestDelegate: server.app,
	}
	httpServer := http.Server{}
	httpServer.Addr = server.config.HttpConfig.Host + ":" + server.config.HttpConfig.Port
	httpServer.Handler = handler
	group, ctx := errgroup.WithContext(server.ctx)
	group.Go(func() error {
		go func() {
			<-ctx.Done()
			_ = httpServer.Shutdown(ctx)
		}()
		return httpServer.ListenAndServe()
		// return http.ListenAndServe(":8080", nil)
	})
	group.Go(func() error {
		<-ctx.Done()
		return errors.New("http server exit due to outer context done")
	})
	return group.Wait()
}

func (server *Server) prepare() {
	route.ScanEndPoint(server.endPointTrie, server.controllers, server.config.HttpConfig.Prefix)
	server.buildPipeline()
}

// buildPipeline will get RequestPipeline ready for request
func (server *Server) buildPipeline() {
	server.app = server.pipeline.Build(server.endPointTrie)
}
