package Route

import "reflect"

type Controller interface {
	GetControllerType() reflect.Type
}
