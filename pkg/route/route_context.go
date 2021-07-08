package route

import (
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	http3 "github.com/wwbweibo/EasyRoute/pkg/http"
	TypeManagement2 "github.com/wwbweibo/EasyRoute/pkg/types"
	"net/http"
	"strings"
)

var typeCollectionInstance = TypeManagement2.NewTypeCollect()

type RouteContext struct {
	controllers []*controllers.Controller // 添加到上下文中的控制器
	// routeMap       map[string]routeMap         // 用于保存终结点和处理方法的映射
	endPointTrie   *EndPointTrie                // 终结点前缀树
	pipeline       Pipeline                     // 请求处理管道
	app            delegates.RequestDelegate    // 最终的请求处理方法
	typeCollection *TypeManagement2.TypeCollect // 内置类型字典
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
func (receiver *RouteContext) AddController(controller controllers.Controller) {
	receiver.controllers = append(receiver.controllers, &controller)
}

func (receiver *RouteContext) AddMiddleware(middleware delegates.Middleware) {
	receiver.pipeline.AddMiddleware(middleware)
}

// add a default action for given pattern
func (receiver *RouteContext) AddDefaultHandler(pattern string, delegate delegates.RequestDelegate) {
	target_node, _ := receiver.endPointTrie.GetRoot().Search(strings.Split(pattern, "/")[1:])
	target_node.defaultHandler = delegate
}

func (receiver *RouteContext) RegisterTypeByInstance(instance interface{}) {
	receiver.typeCollection.Register(instance)
}

func (receiver *RouteContext) HandleRequest(ctx *http3.HttpContext) {
	receiver.app(ctx)
}

func (receiver *RouteContext) Serve() error {
	http.HandleFunc(receiver.endpointPrefix, func(writer http.ResponseWriter, request *http.Request) {
		ctx := http3.HttpContext{
			request,
			writer,
		}
		receiver.app(&ctx)
	})
	return http.ListenAndServe(":8080", nil)

}

func (receiver *RouteContext) buildPipeline() {
	receiver.app = receiver.pipeline.build()
}

// find endpoint from given Controller list
func (receiver *RouteContext) routeParse(prefix string) {
	scanEndPoint(receiver, prefix)
}
