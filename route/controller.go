package route

import "reflect"

type Controller interface {
	GetControllerType() reflect.Type
}
