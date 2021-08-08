package main

import (
	"context"
	"fmt"
	"github.com/wwbweibo/EasyRoute"
	"reflect"
)

func main() {
	// logger.WithLogger(adapter.LogrusAdapter{})
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
	Index       func() (string, error)              `method:"Get"`
	IndexA      func(a string) (string, error)      `method:"Get" param:"a"`
	IndexPerson func(person Person) (Person, error) `method:"get" param:"person"`
	PostIndex   func() (string, error)              `method:"POST"`
	PostPerson  func(person Person) (Person, error) `method:"POST" param:"person:FromForm"`
}

func (self *HomeController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*self)
}

func NewHomeController() *HomeController {
	instance := HomeController{
		Index: func() (string, error) {
			fmt.Println("enter index")
			return "Index", nil
		},
		IndexA: func(a string) (string, error) {
			fmt.Println(a)
			return a, nil
		},
		IndexPerson: func(person Person) (Person, error) {
			fmt.Printf("Name： %s, Age：%f\n", person.Name, person.Age)
			return person, nil
		},
		PostIndex: func() (string, error) {
			return "success", nil
		},
		PostPerson: func(person Person) (Person, error) {
			return person, nil
		},
	}
	return &instance
}
