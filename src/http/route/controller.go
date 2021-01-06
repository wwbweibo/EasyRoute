package route

import "reflect"

// 控制器接口，所有控制器需要实现该接口

type Controller interface {
	GetControllerType() reflect.Type
}
