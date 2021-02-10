package di

import "reflect"

type IContainer interface {
	Resolve(p reflect.Type) interface{}
}
