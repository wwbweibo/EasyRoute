package main

import "github.com/wwbweibo/EasyRoute/src/http"

func main() {
	server := http.NewHttpServer("0.0.0.0", "80")
	server.Serve()
}
