package main

import (
	"reflect"
)

type Person struct {
	Name string  `json:"Name"`
	Age  float64 `json:"Age"`
}

//go:generate rpcgenerator
type HomeController struct {
	Index       func() (string, error)              `method:"Get"`
	IndexA      func(a string) (string, error)      `method:"Get" param:"a"`
	IndexPerson func(person Person) (Person, error) `method:"get" param:"person"`
	PostIndex   func() (string, error)              `method:"POST"`
}

func (controller *HomeController) GetControllerType() reflect.Type {
	panic("implement me")
}
