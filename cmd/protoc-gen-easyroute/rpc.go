package main

import (
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func generateRpcClient(file *protogen.GeneratedFile, service *protogen.Service) {
	file.P(fmt.Sprintf("func New%sClient(config rpc.Config) *%sController {", service.GoName, service.GoName))

	file.P(fmt.Sprintf("\treturn &%sController {", service.GoName))
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		httpMethod := "GET"
		rpcMethod := "HttpGet"
		value := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
		rule := value.(*annotations.HttpRule)
		if rule != nil {
			_, httpMethod = resolveHttpMethod(rule)
		}
		file.P(fmt.Sprintf("\t\t%s: func(ctx context.Context, input %s)(%s, error) {",
			method.GoName,
			method.Input.GoIdent.GoName,
			method.Output.GoIdent.GoName))
		file.P(fmt.Sprintf("\t\t\t result := %s{}", method.Output.GoIdent.GoName))
		switch httpMethod {
		case "GET":
			rpcMethod = "HttpGet"
		}
		file.P(fmt.Sprintf("\t\t\tparsedInput := rpc.ParseInput(input)"))
		file.P(fmt.Sprintf("\t\t\terr := rpc.%s(config,\"%s\", \"%s\", parsedInput, &result)", rpcMethod, service.GoName, method.GoName))
		file.P(fmt.Sprintf("\t\t\treturn result, err"))
		file.P(fmt.Sprintf("\t\t},"))
	}
	file.P("\t}")
	//for _, method := range service.Methods {
	//    if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
	//        continue
	//    }
	//    file.P(fmt.Sprintf("    %s func(ctx context.Context, input %s) (%s, error) `%s`",
	//        method.GoName,
	//        method.Input.GoIdent.GoName,
	//        method.Output.GoIdent.GoName,
	//        getMethodTag(method)))
	//}
	file.P("}")
}
