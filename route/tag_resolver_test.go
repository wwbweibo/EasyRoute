package route

import (
	"reflect"
	"testing"
)

type Person struct {
}

type TestController struct {
	Index          func() (string, error)
	Index1         func() (string, error)             `route:"home" method:"Post"`
	IndexWithParam func(a string, b Person, c Person) `param:"a,person,p:FromForm"`
}

func (t *TestController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*t)
}

type Test1Controller struct {
	controllerName string
}

func (t *Test1Controller) GetControllerType() reflect.Type {
	return reflect.TypeOf(*t)
}

func TestResolveControllerName_NotDefineName(t *testing.T) {
	controllerName := ResolveControllerName(&TestController{})
	if controllerName != "test" {
		t.Errorf("controller name %s is no matched", controllerName)
		t.Fail()
	}
}

func TestResolveControllerName_DefinedName(t *testing.T) {
	controllerName := ResolveControllerName(&Test1Controller{controllerName: "ABC"})
	if controllerName != "abc" {
		t.Errorf("controller name %s is no matched", controllerName)
		t.Fail()
	}
}

func TestResolveMethodName(t *testing.T) {
	controller := TestController{}
	controllerType := controller.GetControllerType()
	method1 := ResolveMethodName(controllerType.Field(0))
	if method1 != "index" {
		t.Errorf("method name %s is not matched", method1)
		t.Fail()
	}
	method2 := ResolveMethodName(controllerType.Field(1))
	if method2 != "home" {
		t.Errorf("method name %s is not matched", method2)
		t.Fail()
	}
}

func TestResolveMethod(t *testing.T) {
	controller := TestController{}
	controllerType := controller.GetControllerType()
	method1 := ResolveMethod(controllerType.Field(0).Tag)
	if method1 != "GET" {
		t.Errorf("method %s is not matched", method1)
		t.Fail()
	}
	method2 := ResolveMethod(controllerType.Field(1).Tag)
	if method2 != "POST" {
		t.Errorf("method %s is not matched", method2)
		t.Fail()
	}
}
