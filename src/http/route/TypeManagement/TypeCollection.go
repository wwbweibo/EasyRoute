package TypeManagement

import (
	"errors"
	"reflect"
)

var internalTypes = []reflect.Type{
	reflect.TypeOf(int(1)),
	reflect.TypeOf(int8(1)),
	reflect.TypeOf(int16(1)),
	reflect.TypeOf(int32(1)),
	reflect.TypeOf(int64(1)),
	reflect.TypeOf(float32(1)),
	reflect.TypeOf(float64(1)),
	reflect.TypeOf(true),
	reflect.TypeOf(""),
}

type TypeCollect struct {
	types map[string]*reflect.Type
}

func NewTypeCollect() *TypeCollect {
	instance := &TypeCollect{
		types: make(map[string]*reflect.Type),
	}
	instance.registerInternalTypes()
	return instance
}

func (receiver *TypeCollect) registerInternalTypes() {
	for _, internalType := range internalTypes {
		receiver.types[internalType.String()] = &internalType
	}
}

func (receiver *TypeCollect) Register(inst interface{}) {
	instType := reflect.TypeOf(inst)
	receiver.RegisterType(&instType)
}

func (receiver *TypeCollect) RegisterType(t *reflect.Type) {
	x := *t

	if x.Kind() == reflect.Ptr {
		x = x.Elem()
	}

	if _, ok := receiver.types[x.String()]; ok {
		return
	} else {
		receiver.types[x.String()] = &x
	}
}

func (receiver *TypeCollect) TypeOf(name string) (*reflect.Type, error) {
	if t, ok := receiver.types[name]; ok {
		return t, nil
	} else {
		return nil, errors.New("Error to get type of :" + name)
	}
}

func (receiver *TypeCollect) InstanceOf(name string) (reflect.Value, error) {
	if t, ok := receiver.types[name]; ok {
		createdInstance := reflect.New(*t)
		a := createdInstance.Elem()
		return a, nil
		// return createdInstance.Elem().Interface(), nil
	} else {
		return reflect.Value{}, errors.New("Could not find type: " + name + ", Please sure you have registered first")
	}
}
