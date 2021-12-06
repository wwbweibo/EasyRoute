# protoc-gen-easyroute

this is a grpc plugin used to generate easyroute code from proto files

currently, it only supports the following features:
- automatic generation of the `Controller` struct and implementation of the `Controller` interface
- generate rpc client code

to use this tools, please install the following dependencies:
- protoc
- protoc-gen-go

to install the plugin, please run the following command:

```bash
go install github.com/wwbweibo/EasyRoute/protoc-gen-easyroute@latest
```

to generate controller code from proto files, you can use the following command:
```bash
protoc -I [third_party] --proto_path [directory of proto file] --go_out . --easyroute_out . [proto file] 
```
if you want to generate the rpc client code, you can use the following command:
```bash
protoc -I [third_party] --proto_path [directory of proto file] --go_out . --easyroute_out=rpc=true:. [proto file] 
```
this command will generate the rpc client code in the same file as the controller code. 