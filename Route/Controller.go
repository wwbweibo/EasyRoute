package Route

import "reflect"

type Controller interface {
	RegisterAsController(ctx *RouteContext)
	GetControllerType() reflect.Type
}
