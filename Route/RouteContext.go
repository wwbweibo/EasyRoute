package Route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type RouteContext struct {
	controllers []*Controller
	routeMap    map[string]routeMap
}

type routeMap struct {
	endPoint       string
	method         string
	controller     *Controller
	controllerType reflect.Type
	controllerName string
}

func NewRouteContext() *RouteContext {
	return &RouteContext{
		controllers: make([]*Controller, 0),
	}
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
	path := c.Request.RequestURI

	if routeMap, ok := self.routeMap[path]; ok {
		if c.Request.Method == routeMap.method {
			methodName := strings.Replace(path, "/"+routeMap.controllerName+"/", "", 1)
			result := reflect.ValueOf(*routeMap.controller).Elem().FieldByName(methodName).Call(nil)[0]
			c.String(http.StatusOK, result.String())
		} else {
			c.String(http.StatusMethodNotAllowed, "405 NotAllowed")
		}

	} else {
		c.String(http.StatusNotFound, "404 NotFind")
	}
}
