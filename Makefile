.PHONY: compile
compile:
	buf generate

PHONY: lint
lint:
	buf lint

.PHONY: format
format:
	buf format

.PHONY: gen
	go mod init github.com/almalii/grpc-contracts/gen/go/google && go mod tidy
	go mod init github.com/almalii/grpc-contracts/gen/go/users_service && go mod tidy
	go mod init github.com/almalii/grpc-contracts/gen/go/auth_service && go mod tidy
	go mod init github.com/almalii/grpc-contracts/gen/go/notes_service && go mod tidy