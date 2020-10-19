package Route

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

var routeContext = RouteContext{
	controllers: make([]*Controller, 0),
	pipeline: Pipeline{
		handlerList: make([]Middleware, 0),
	},
}

type RouteContext struct {
	controllers []*Controller       // 添加到上下文中的控制器
	routeMap    map[string]routeMap // 用于保存终结点和处理方法的映射
	pipeline    Pipeline            // 请求处理管道
	app         RequestDelegate     // 最终的请求处理方法
}

type HttpContext *gin.Context

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

// Init the RouteContext and begin listen
func (self *RouteContext) InitRouteWithGivenController(controllers []*Controller, listenAddr string) {
	self.controllers = controllers
	self.RouteParse()
	self.Start(listenAddr)
}

// Init the RouteContext and begin listen
func (self *RouteContext) InitRoute(listenAddr string) {
	self.RouteParse()
	self.Start(listenAddr)
}

// add Controller to RouteContext
func (self *RouteContext) AddController(controller Controller) {
	self.controllers = append(self.controllers, &controller)
}

// find endpoint from given Controller list
func (self *RouteContext) RouteParse() {
	set := make(map[string]routeMap)
	for _, controller := range self.controllers {
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
	self.routeMap = set
}

func (receiver *RouteContext) AddMiddleware(middleware Middleware) {
	receiver.pipeline.AddMiddleware(middleware)
}

// start http listen using gin
func (self *RouteContext) Start(addr string) {
	self.app = self.pipeline.build()
	router := gin.Default()
	rootGroup := router.Group("/*path")
	rootGroup.Any("", self.route)
	router.Run(addr)
}

func (self *RouteContext) route(c *gin.Context) {
	ctx := HttpContext(c)
	self.app(ctx)
}
