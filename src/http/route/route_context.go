package route

import (
	"github.com/wwbweibo/EasyRoute/src/http"
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"github.com/wwbweibo/EasyRoute/src/http/route/TypeManagement"
	"strings"
)

const (
	INTERNAL_SERVER = "INTERNAL"
	GIN             = "GIN"
)

var routeContext = RouteContext{
	controllers: make([]*Controller, 0),
	pipeline: Pipeline{
		handlerList: make([]Middleware, 0),
	},
	typeCollection: typeCollectionInstance,
	endPointTrie:   NewEndPointTrie(),
}

var typeCollectionInstance = TypeManagement.NewTypeCollect()

type RouteContext struct {
	controllers []*Controller // 添加到上下文中的控制器
	// routeMap       map[string]routeMap         // 用于保存终结点和处理方法的映射
	endPointTrie   *EndPointTrie               // 终结点前缀树
	pipeline       Pipeline                    // 请求处理管道
	app            http.RequestDelegate        // 最终的请求处理方法
	typeCollection *TypeManagement.TypeCollect // 内置类型字典
	endpointPrefix string
	server         interface{}
	serverType     string
	listenPort     string
	listenAddress  string
}

type paramMap struct {
	paramName string
	paramType string
	source    string
}

func NewRouteContext() *RouteContext {
	return &routeContext
}

// 初始化 RouteContext 已准备进行处理
func (receiver *RouteContext) InitRoute(prefix string) {
	receiver.endpointPrefix = prefix
	receiver.routeParse(prefix)
	receiver.buildPipeline()
}

// add Controller to RouteContext
func (receiver *RouteContext) AddController(controller Controller) {
	receiver.controllers = append(receiver.controllers, &controller)
}

func (receiver *RouteContext) AddMiddleware(middleware Middleware) {
	receiver.pipeline.AddMiddleware(middleware)
}

// add a default action for given pattern
func (receiver *RouteContext) AddDefaultHandler(pattern string, delegate http.RequestDelegate) {
	target_node, _ := receiver.endPointTrie.GetRoot().Search(strings.Split(pattern, "/")[1:])
	target_node.defaultHandler = delegate
}

func (receiver *RouteContext) RegisterTypeByInstance(instance interface{}) {
	receiver.typeCollection.Register(instance)
}

func (receiver *RouteContext) HandleRequest(ctx *context.Context) {
	receiver.app(ctx)
}

func (receiver *RouteContext) WithServer(serverName string, host string, port string) *RouteContext {
	if serverName == "" || serverName == INTERNAL_SERVER {
		server := http.NewHttpServer(host, port, receiver.app)
		receiver.server = server
		receiver.serverType = INTERNAL_SERVER
		receiver.listenPort = port
		receiver.listenAddress = host
	}

	return receiver
}

func (receiver *RouteContext) Serve() error {
	if receiver.serverType == "" || receiver.serverType == INTERNAL_SERVER {
		server := receiver.server.(*http.Server)
		return server.Serve()
	}
	return nil
}

func (receiver *RouteContext) buildPipeline() {
	receiver.app = receiver.pipeline.build()
}

// find endpoint from given Controller list
func (receiver *RouteContext) routeParse(prefix string) {
	scanEndPoint(receiver, prefix)
}
