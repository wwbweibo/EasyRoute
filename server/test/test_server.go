package main

import server2 "github.com/wwbweibo/EasyRoute/server"

func main() {
	server := server2.NewServer("0.0.0.0", "80")
	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
