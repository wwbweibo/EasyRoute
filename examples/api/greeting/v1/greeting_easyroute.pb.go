// Code generated by protoc-gen-easyroute. DO NOT EDIT.
// versions:
// protoc-gen-easyroute 1.0.0

package v1

import (
	context "context"
	codec "github.com/lazada/protoc-gen-go-http/codec"
	http "github.com/wwbweibo/EasyRoute/http"
	rpc "github.com/wwbweibo/EasyRoute/rpc"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

type GreetingServiceController struct {
	server   GreetingServiceServer
	Greeting func(ctx *http.Context) `method:"POST" route:"/v1/greeting" param:"input"`
}

func NewGreetingServiceController(server GreetingServiceServer) *GreetingServiceController {
	return &GreetingServiceController{
		server: server,
		Greeting: func(ctx *http.Context) {
			input := &GreetingRequest{}
			codec.NewJsonRPCCodec().ReadRequest(ctx.Request, input)
			result, err := server.Greeting(ctx.Ctx, input)
			if err != nil {
				ctx.WriteError(err)
				return
			}
			ctx.WriteJson(result, 200)
		},
	}
}

func (c *GreetingServiceController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*c)
}

type GreetingServiceHttpClient struct {
	config rpc.Config
}

func NewGreetingServiceHttpClient(config rpc.Config) *GreetingServiceHttpClient {
	return &GreetingServiceHttpClient{config: config}
}
func (client *GreetingServiceHttpClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	result := GreetingResponse{}
	parsedInput := rpc.ParseInput(in)
	err := rpc.HttpGet(client.config, "GreetingService", "Greeting", parsedInput, &result)
	return &result, err
}
