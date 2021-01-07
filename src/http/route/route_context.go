package route

import (
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"github.com/wwbweibo/EasyRoute/src/http/route/TypeManagement"
	"reflect"
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
	app            RequestDelegate             // 最终的请求处理方法
	typeCollection *TypeManagement.TypeCollect // 内置类型字典
}

type routeMap struct {
	endPoint       string       // 保存终结点信息
	method         string       // 保存请求方法信息
	controller     *Controller  // 保存对应的控制器信息
	controllerType reflect.Type // 保存控制器类型信息
	controllerName string       // 控制器名称
	paramMap       *[]paramMap  // 参数来源
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
func (receiver *RouteContext) InitRoute() {
	receiver.RouteParse()
	receiver.buildPipeline()
}

// add Controller to RouteContext
func (receiver *RouteContext) AddController(controller Controller) {
	receiver.controllers = append(receiver.controllers, &controller)
}

// find endpoint from given Controller list
func (receiver *RouteContext) RouteParse() {
	scanEndPoint(receiver)
}

func (receiver *RouteContext) AddMiddleware(middleware Middleware) {
	receiver.pipeline.AddMiddleware(middleware)
}

func (receiver *RouteContext) RegisterTypeByInstance(instance interface{}) {
	receiver.typeCollection.Register(instance)
}

func (receiver *RouteContext) HandleRequest(ctx *context.Context) {
	receiver.app(ctx)
}

func (receiver *RouteContext) buildPipeline() {
	receiver.app = receiver.pipeline.build()
}
