.PHONY: compile

compile:
	protoc --go_out=. --go-grpc_out=. ./auth_service/model/v1/auth.proto
	protoc --go_out=. --go-grpc_out=. ./auth_service/service/v1/auth.proto

.PHONY: gen
	go mod init GoProjects/grpc-contracts/gen/go/auth_service
