package EasyRoute

import (
	"context"
	"errors"
	controllers2 "github.com/wwbweibo/EasyRoute/controllers"
	delegates2 "github.com/wwbweibo/EasyRoute/delegates"
	tm "github.com/wwbweibo/EasyRoute/internal/types"
	"github.com/wwbweibo/EasyRoute/logger"
	route2 "github.com/wwbweibo/EasyRoute/route"
	"golang.org/x/sync/errgroup"
	"net/http"
	"reflect"
	"strings"
)

type Server struct {
	config       Config
	ctx          context.Context
	controllers  []controllers2.Controller  // controllers is all the Controller add to current server
	pipeline     *route2.Pipeline           // pipeline defines the request pipeline for current server
	app          delegates2.RequestDelegate // app is the entry point for a request
	endPointTrie *route2.EndPointTrie       // endPointTrie maintain all resolved endpoint
	ts           *tm.TypeCollect            // tsn maintain all registered type used in controllers
}

func NewServer(ctx context.Context, config Config) (*Server, error) {
	// routeContext := route.NewRouteContext(ctx)
	return &Server{
		config:       config,
		ctx:          ctx,
		controllers:  []controllers2.Controller{},
		pipeline:     route2.NewPipeline(),
		endPointTrie: route2.NewEndPointTrie(),
		ts:           tm.NewTypeCollect(),
	}, nil
}

// AddController will register the given handle to server
func (server *Server) AddController(controller controllers2.Controller) {
	logger.Info("[server] - [AddController] add controller " + controller.GetControllerType().String())
	server.controllers = append(server.controllers, controller)
}

// AddMiddleware will add a middle to pipeline
func (server *Server) AddMiddleware(middleware delegates2.Middleware) {
	logger.Info("[server] - [AddMiddleware] add middleware " + reflect.TypeOf(middleware).String())
	server.pipeline.AddMiddleware(middleware)
}

// AddDefaultHandler add a default action for given pattern
func (server *Server) AddDefaultHandler(pattern string, delegate delegates2.RequestDelegate) {
	logger.Info("[server] - [AddDefaultHandler] add default handler to pattern " + pattern)
	targetNode, _ := server.endPointTrie.GetRoot().Search(strings.Split(pattern, "/")[1:])
	targetNode.DefaultHandler = delegate
}

// RegisterType will register given type server type collection
func (server *Server) RegisterType(t interface{}) {
	server.ts.Register(t)
}

// Serve will make server ready for request and listen with given config
func (server *Server) Serve() error {
	server.prepare()
	handler := &route2.HttpRequestHandler{
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
			logger.Info("[server] - [Serve] server shutdown due to outer context exit")
		}()
		logger.Info("[server] - [Serve] listen http on address " + httpServer.Addr)
		return httpServer.ListenAndServe()
	})
	group.Go(func() error {
		<-ctx.Done()
		return errors.New("http server exit due to outer context done")
	})
	return group.Wait()
}

func (server *Server) prepare() {
	route2.InjectTypes(server.ts)
	route2.ScanEndPoint(server.endPointTrie, server.controllers, server.config.HttpConfig.Prefix)
	server.buildPipeline()
}

// buildPipeline will get RequestPipeline ready for request
func (server *Server) buildPipeline() {
	server.app = server.pipeline.Build(server.endPointTrie)
}
