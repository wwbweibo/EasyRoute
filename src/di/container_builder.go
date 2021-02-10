package di

import (
	"reflect"
)

type IContainerBuilder interface {
	Build() IContainer
	AddInstance(inst interface{})
	AddType(t reflect.Type)
	AddTypes(t []reflect.Type)
}

type DefaultContainerBuilder struct {
	registeredType     map[reflect.Type]bool
	registeredInstance map[reflect.Type]interface{}
}

func NewDefaultContainerBuilder() IContainerBuilder {
	return &DefaultContainerBuilder{
		registeredType:     make(map[reflect.Type]bool),
		registeredInstance: make(map[reflect.Type]interface{}),
	}
}

func (d DefaultContainerBuilder) Build() IContainer {
	// todo 构建Container的方法
	panic("implement me")
}

func (d DefaultContainerBuilder) AddInstance(inst interface{}) {
	instanceType := reflect.TypeOf(inst)
	for instanceType.Kind() == reflect.Ptr {
		instanceType = instanceType.Elem()
	}
	d.registeredInstance[instanceType] = inst
}

func (d DefaultContainerBuilder) AddType(t reflect.Type) {
	d.registeredType[t] = true
}

func (d DefaultContainerBuilder) AddTypes(t []reflect.Type) {
	for _, t := range t {
		d.AddType(t)
	}
}
