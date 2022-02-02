package main

import (
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func generateController(file *protogen.GeneratedFile, service *protogen.Service) {
	file.P(fmt.Sprintf("type %sController struct {", service.GoName))
	file.P(fmt.Sprintf("\tserver %sServer", service.GoName))
	// writing GetControllerType function
	// file.P(fmt.Sprintf("    GetControllerType() reflect.Type"))
	// walk through all the service method and generate function
	for _, method := range service.Methods {
		if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
			continue
		}
		file.P(fmt.Sprintf("    %s func(ctx *http.Context) `%s`",
			method.GoName,
			getMethodTag(method)))
	}
	file.P("}\n\n")
	generateConstructor(file, service)
	file.P("\n\n")
	generateInterfaceImplement(file, service)
	if *enableRpc {
		generateRpcClient(file, service)
	}
}

// getMethodTag will accept the method and generate the tags for the method
func getMethodTag(method *protogen.Method) string {
	value := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
	rule := value.(*annotations.HttpRule)
	if rule != nil {
		path, method := resolveHttpMethod(rule)
		// TODO: 解析参数
		return fmt.Sprintf("method:\"%s\" route:\"%s\" param:\"input\"", method, path)
	}
	return fmt.Sprintf("method:\"GET\"")
}

func resolveHttpMethod(rule *annotations.HttpRule) (string, string) {
	var path string
	var method string
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = "GET"
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = "PUT"
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = "POST"
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = "DELETE"
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = "PATCH"
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}
	return path, method
}

func generateConstructor(file *protogen.GeneratedFile, service *protogen.Service) {
	file.P(fmt.Sprintf("func New%sController(server %sServer) *%sController {",
		service.GoName,
		service.GoName,
		service.GoName))
	file.P(fmt.Sprintf("\treturn &%sController{", service.GoName))
	file.P(fmt.Sprintf("\t\tserver: server,"))
	for _, method := range service.Methods {
		generateControllerMethodImpl(file, service, method)
	}
	file.P("\t}")
	file.P("}")
}

func generateControllerMethodImpl(file *protogen.GeneratedFile, service *protogen.Service, method *protogen.Method) {
	file.P(fmt.Sprintf("\t\t%s: func(ctx *http.Context) {", method.GoName))
	file.P(fmt.Sprintf("\t\t\tinput := &%s{}", method.Input.GoIdent.GoName))
	file.P("\t\t\tcodec.NewJsonRPCCodec().ReadRequest(ctx.Request, input)")
	file.P(fmt.Sprintf("\t\t\tresult, err := server.%s(ctx.Ctx, input)", method.GoName))
	file.P("\t\t\tif err != nil {")
	file.P("\t\t\t\t ctx.WriteError(err)")
	file.P("\t\t\t\t return")
	file.P("\t\t\t}")
	file.P("\t\t\tcodec.NewJsonRPCCodec().WriteResponse(ctx.Response, result)")
	file.P("\t\t},")
}

func generateInterfaceImplement(file *protogen.GeneratedFile, service *protogen.Service) {
	file.P(fmt.Sprintf("func (c *%sController) GetControllerType() reflect.Type {", service.GoName))
	file.P(fmt.Sprintf("    return reflect.TypeOf(*c)"))
	file.P("}")
}
