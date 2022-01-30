package main

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/rpc"
)

func main() {
	config := rpc.Config{BaseUrl: "http://localhost:8080"}
	controller := NewHomeController(config)
	result, err := controller.Index()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
