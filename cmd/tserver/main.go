package main

import (
	"context"
	"fmt"
	"github.com/wwbweibo/EasyRoute"
	"github.com/wwbweibo/EasyRoute/logger"
	"github.com/wwbweibo/EasyRoute/logger/adapter"
	"reflect"
)

func main() {
	logger.WithLogger(adapter.LogrusAdapter{})
	ctx, _ := context.WithCancel(context.Background())
	config := EasyRoute.Config{
		HttpConfig: EasyRoute.HttpConfig{
			Prefix: "/",
			Host:   "0.0.0.0",
			Port:   "8080",
		},
	}
	server, _ := EasyRoute.NewServer(ctx, config)
	server.RegisterType(Person{})
	server.AddController(NewHomeController())
	err := server.Serve()
	fmt.Println(err.Error())
}

type Person struct {
	Name string  `json:"Name"`
	Age  float64 `json:"Age"`
}

type HomeController struct {
	Index       func() string              `method:"Get"`
	IndexA      func(a string) string      `method:"Get" param:"a"`
	IndexPerson func(person Person) Person `method:"get" param:"person"`
}

func (self *HomeController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*self)
}

func NewHomeController() *HomeController {
	instance := HomeController{
		Index: func() string {
			fmt.Println("enter index")
			return "Index"
		},
		IndexA: func(a string) string {
			fmt.Println(a)
			return a
		},
		IndexPerson: func(person Person) Person {
			fmt.Printf("Name： %s, Age：%f\n", person.Name, person.Age)
			return person
		},
	}
	return &instance
}
