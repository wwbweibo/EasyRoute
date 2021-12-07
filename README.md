# EasyRoute

EasyRoute provide an easy way to create http server and expose your end point.

## controller

EasyRoute enables you to use controller in golang.  
To create a controller, just create a struct witch name end with `Controller`, and then implement `EasyRoute.controllers.Controller` interface. 

```go
type HomeController struct {
    Index       func(ctx *http.Context)             `method:"GET"`
}
```
In this example, the server will scan this struct, and resolve the controller name as `Home`, if you want to customs controller name just provide a string field named `controllerName`. In this way, you need set the field value in your constructor.  

For the function, you need to provide a list of tag to tell server how to expose it. Current, EasyRoute allows you to use tags as follows:

- method: the http request method for the function, if not provide, `GET` as default.
- route: the exposed name for the function, optional, default is function name.

the server will expose the function as `/[ControllerName]/[route]`

please note that the function name must be `func(ctx *http.Context)`

## initial and start a server

to initial a server, just create a server instance, add your controller and type used in your controller, and then, call `server.Serve` to start server, you can find an example in `cmd/tserver`

```go
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
```

please note, you must call `server.RegisterType` to register your type used in your controller, since golang does not provide any function to obtain type information or create type instance using TypeName. 

## rpc

EasyRoute allow you to use rpc for controller. EasyRoute to code generate to create client rpc code.  
first, you need add rpcgenerator to you PATH, to do this, clone the repo, cd `cmd/rpcgenerator`, and execute `go install`  
then, add `//go:generate rpcgenerator` comment to controller.  
last, execute `go generate` to generate code.

## protobuf

this library is also support protobuf, you can use proto file to define your route.  

about how to use proto file to define your route, please refer to [protoc-gen-easyroute](cmd/protoc-gen-easyroute/readme.md)