package Route

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

type RouteContext struct {
	pipeline    []PipelineHandler
	controllers []*Controller
	routeMap    map[string]routeMap
}

type RequestContext *gin.Context

type routeMap struct {
	endPoint       string
	method         string
	controller     *Controller
	controllerType reflect.Type
	controllerName string
}

func NewRouteContext() *RouteContext {
	inst := RouteContext{
		controllers: make([]*Controller, 0),
		pipeline:    []PipelineHandler{},
	}
	reqHandler := requestHandler{
		routeContext: &inst,
	}

	inst.AddPipeline(&reqHandler)
	return &inst
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

func (self *RouteContext) AddPipeline(pipeline PipelineHandler) {
	self.pipeline = append(self.pipeline, pipeline)
}

// find endpoint from given Controller list
func (self *RouteContext) RouteParse() {
	set := make(map[string]routeMap)
	for _, controller := range self.controllers {
		controllerType := (*controller).GetControllerType()
		patharr := strings.Split(controllerType.String(), ".")
		controllerName := strings.Replace(patharr[len(patharr)-1], "Controller", "", 1)
		for i := 0; i < controllerType.NumField(); i++ {
			field := controllerType.Field(i)
			route := field.Tag.Get("Route")
			method := field.Tag.Get("method")
			if strings.Contains(route, "{Controller}") {
				route = strings.Replace(route, "{Controller}", controllerName, 1)
			}
			set[route] = routeMap{
				endPoint:       route,
				method:         strings.ToUpper(method),
				controller:     controller,
				controllerType: controllerType,
				controllerName: controllerName,
			}
		}
	}
	self.routeMap = set
}

// start http listen using gin
func (self *RouteContext) Start(addr string) {
	router := gin.Default()
	rootGroup := router.Group("/*path")
	rootGroup.Any("", self.route)
	router.Run(addr)
}

func (self *RouteContext) route(c *gin.Context) {
	context := RequestContext(c)
	for i := len(self.pipeline) - 1; i >= 1; i-- {
		self.pipeline[i].Handle(context, self.pipeline[i-1])
	}
	self.pipeline[0].Handle(c, nil)
}
