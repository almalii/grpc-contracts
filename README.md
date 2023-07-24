#### Commands to generate gRPC code from proto files
- `protoc --go_out=../. --go-grpc_out=../. ./auth.proto`
- `protoc --go_out=../. --go-grpc_out=../. ./notes.proto`
- `protoc --go_out=../. --go-grpc_out=../. ./users.proto` 

`--go_out` *is used to generate Go code*

`--go-grpc_out` *is used to generate gRPC code*

