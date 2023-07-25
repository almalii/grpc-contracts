.PHONY: compile

compile:
	buf generate
.PHONY: gen
	go mod init github.com/almalii/grpc-contracts/gen/go/auth_service && go mod tidy
