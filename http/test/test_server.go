package main

import "github.com/wwbweibo/EasyRoute/http"

func main() {
	server := http.NewHttpServer("0.0.0.0", "80")
	server.Serve()
}
