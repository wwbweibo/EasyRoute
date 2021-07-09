package route

import (
	"github.com/wwbweibo/EasyRoute/pkg/controllers"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	iHttp "github.com/wwbweibo/EasyRoute/pkg/http"
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

// InitRoute init RouteContext ready for request
func (context *RouteContext) InitRoute(prefix string) {
	context.endpointPrefix = prefix
	context.routeParse(prefix)
	context.buildPipeline()
}

// add Controller to RouteContext
func (context *RouteContext) AddController(controller controllers.Controller) {
	context.controllers = append(context.controllers, &controller)
}

func (context *RouteContext) AddMiddleware(middleware delegates.Middleware) {
	context.pipeline.AddMiddleware(middleware)
}

// add a default action for given pattern
func (context *RouteContext) AddDefaultHandler(pattern string, delegate delegates.RequestDelegate) {
	target_node, _ := context.endPointTrie.GetRoot().Search(strings.Split(pattern, "/")[1:])
	target_node.defaultHandler = delegate
}

func (context *RouteContext) RegisterTypeByInstance(instance interface{}) {
	context.typeCollection.Register(instance)
}

func (context *RouteContext) HandleRequest(ctx *iHttp.HttpContext) {
	context.app(ctx)
}

func (context *RouteContext) Serve() error {
	http.HandleFunc(context.endpointPrefix, func(writer http.ResponseWriter, request *http.Request) {
		ctx := iHttp.HttpContext{
			request,
			writer,
		}
		context.app(&ctx)
	})
	return http.ListenAndServe(":8080", nil)

}

func (context *RouteContext) buildPipeline() {
	context.app = context.pipeline.build()
}

// find endpoint from given Controller list
func (context *RouteContext) routeParse(prefix string) {
	scanEndPoint(context, prefix)
}
