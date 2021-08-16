Commands to install protoc:

```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

Commands to generate go from proto:

```shell
protoc --go_out=plugins=grpc:chat  --go_opt=paths=source_relative chat.proto
```