package main

import (
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func generateRpcClientStruct(file *protogen.GeneratedFile, service *protogen.Service) {
	file.P(fmt.Sprintf("type %sHttpClient struct {", service.GoName))
	file.P(fmt.Sprintf("\tconfig rpc.Config"))
	file.P(fmt.Sprintf("}"))
}

func generateRpcClientConstructor(file *protogen.GeneratedFile, service *protogen.Service) {
	file.P(fmt.Sprintf("func New%sHttpClient(config rpc.Config) *%sHttpClient {", service.GoName, service.GoName))
	file.P(fmt.Sprintf("\treturn &%sHttpClient{config: config}", service.GoName))
	file.P("}")
}

func generateRpcClient(file *protogen.GeneratedFile, service *protogen.Service) {
	file.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/wwbweibo/EasyRoute/rpc"})
	generateRpcClientStruct(file, service)
	generateRpcClientConstructor(file, service)
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		generateMethodImpl(file, service, method)
	}
}

func generateMethodImpl(file *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	httpMethod := "GET"
	rpcMethod := "HttpGet"
	value := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
	rule := value.(*annotations.HttpRule)
	if rule != nil {
		_, httpMethod = resolveHttpMethod(rule)
	}
	file.P(fmt.Sprintf("func (client *%sHttpClient) %s(ctx context.Context, in *%s, opts ...grpc.CallOption) (*%s, error) {",
		service.GoName,
		method.GoName,
		method.Input.GoIdent.GoName,
		method.Output.GoIdent.GoName))
	file.P(fmt.Sprintf("\t result := %s{}", method.Output.GoIdent.GoName))
	switch httpMethod {
	case "GET":
		rpcMethod = "HttpGet"
	}
	file.P(fmt.Sprintf("\tparsedInput := rpc.ParseInput(in)"))
	file.P(fmt.Sprintf("\terr := rpc.%s(client.config,\"%s\", \"%s\", parsedInput, &result)", rpcMethod, service.GoName, method.GoName))
	file.P(fmt.Sprintf("\treturn &result, err"))
	file.P("}")
}
