package main

import (
	"context"
	"fmt"
	"github.com/wwbweibo/EasyRoute"
	v1 "github.com/wwbweibo/EasyRoute/examples/api/greeting/v1"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	config := EasyRoute.Config{
		HttpConfig: EasyRoute.HttpConfig{
			Prefix: "/",
			Host:   "0.0.0.0",
			Port:   "8081",
		},
	}
	server, _ := EasyRoute.NewServer(ctx, config)
	server.AddController(v1.NewGreetingServiceController(&GreetingService{}))
	err := server.Serve()
	fmt.Println(err.Error())
}

type GreetingService struct {
	v1.UnimplementedGreetingServiceServer
}

func (svc *GreetingService) Greeting(ctx context.Context, req *v1.GreetingRequest) (*v1.GreetingResponse, error) {
	return &v1.GreetingResponse{
		Message: "hello " + req.Name,
	}, nil
}
