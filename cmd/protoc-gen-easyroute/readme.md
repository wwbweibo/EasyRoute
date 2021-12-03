# protoc-gen-easyroute

this is a grpc plugin used to generate easyroute code from proto files

currently, it only supports the following features:
- automatic generation of the `Controller` struct and implementation of the `Controller` interface

to generate controller code from proto files, you can use the following command:
```bash
protoc -I [third_party] --proto_path [directory of proto file] --go_out . --easyroute_out . [proto file] 
```