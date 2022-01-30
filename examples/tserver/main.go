package main

import (
	"context"
	"fmt"
	"github.com/wwbweibo/EasyRoute"
	"github.com/wwbweibo/EasyRoute/http"
	"reflect"
)

func main() {
	// logger.WithLogger(adapter.LogrusAdapter{})
	ctx, _ := context.WithCancel(context.Background())
	config := EasyRoute.Config{
		HttpConfig: EasyRoute.HttpConfig{
			Prefix: "/",
			Host:   "0.0.0.0",
			Port:   "8081",
		},
	}
	server, _ := EasyRoute.NewServer(ctx, config)
	server.AddController(NewSampleController())
	err := server.Serve()
	fmt.Println(err.Error())
}

type Person struct {
	Name string  `json:"Name"`
	Age  float64 `json:"Age"`
}

type SampleController struct {
	Index func(ctx *http.Context)
}

func (c *SampleController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*c)
}

func NewSampleController() *SampleController {
	return &SampleController{
		Index: func(ctx *http.Context) {
			name := ctx.Request.URL.Query()["name"]
			fmt.Println(name[0])
			ctx.Response.WriteHeader(200)
			ctx.Response.Write([]byte(name[0]))
		},
	}
}
