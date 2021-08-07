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

func TestResolveParamName(t *testing.T) {
	controller := TestController{}
	controllerType := controller.GetControllerType()
	params := ResolveParamName(controllerType.Field(2))
	if params[0].paramName != "a" || params[0].paramType != "string" {
		t.Errorf("param %s and type %s is not matched",
			params[0].paramName, params[0].paramType)
		t.Fail()
	}

	if params[1].paramName != "person" || params[1].paramType != "route.Person" {
		t.Errorf("param %s and type %s is not matched",
			params[1].paramName, params[1].paramType)
		t.Fail()
	}

	if params[2].paramName != "p" || params[2].paramType != "route.Person" || params[2].source != "FromForm" {
		t.Errorf("param %s, type %s and source %s is not matched",
			params[2].paramName, params[2].paramType, params[2].source)
		t.Fail()
	}
}
