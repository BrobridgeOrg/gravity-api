# gravity-api

Gravity API definitions which includes `.proto` and `pb.go` files

# Update protobuf definitions

Change working directory then run command in the following:

```
protoc --go_out=plugins=grpc:. *.proto
```
