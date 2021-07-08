package main

import (
	"context"
	"fmt"
	"github.com/wwbweibo/EasyRoute/pkg"
	"reflect"
)

func main() {
	ctx := context.Background()
	server, _ := pkg.NewServer(ctx)
	server.AddController(NewHomeController())
	fmt.Println(server.Serve())
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
