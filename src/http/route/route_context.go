package route

import (
	"github.com/wwbweibo/EasyRoute/src/http/context"
	"github.com/wwbweibo/EasyRoute/src/http/route/TypeManagement"
	"reflect"
	"strings"
)

var routeContext = RouteContext{
	controllers: make([]*Controller, 0),
	pipeline: Pipeline{
		handlerList: make([]Middleware, 0),
	},
	typeCollection: typeCollectionInstance,
}

var typeCollectionInstance = TypeManagement.NewTypeCollect()

type RouteContext struct {
	controllers    []*Controller               // 添加到上下文中的控制器
	routeMap       map[string]routeMap         // 用于保存终结点和处理方法的映射
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
	set := make(map[string]routeMap)
	for _, controller := range receiver.controllers {
		controllerType := (*controller).GetControllerType()
		controllerName := resolveControllerName(&controllerType, controller)
		for i := 0; i < controllerType.NumField(); i++ {
			field := controllerType.Field(i)
			route := resolveMethodName(&field.Tag, &field)
			method := resolveMethod(&field.Tag)
			paramList := resolveParamName(&field.Tag, &field)

			// if the route is not start with "/", then combine the controllerName and route
			if !strings.HasPrefix(route, "/") {
				route = "/" + controllerName + "/" + route
			}

			set[route] = routeMap{
				endPoint:       route,
				method:         strings.ToUpper(method),
				controller:     controller,
				controllerType: controllerType,
				controllerName: controllerName,
				paramMap:       paramList,
			}
		}
	}
	receiver.routeMap = set
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

//// start http listen using gin
//func (self *RouteContext) Start(addr, port string) {
//	self.app = self.pipeline.build()
//	server := http.NewHttpServer(addr, port)
//	server.RegisterHandlers(self)
//	server.Serve()
//}
//
//func (self *RouteContext) route(c *http.Context) {
//	//self.app(ctx)
//}
