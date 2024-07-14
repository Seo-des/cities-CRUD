init:
  go mod init unsia

gen:
  protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

server:
  go run server.go

.PHONY: init gen server