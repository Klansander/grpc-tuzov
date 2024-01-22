



go-run:
	ls cmd/sso
	go run cmd/sso/main.go --config=./config/config.local.yaml
proto_gen:
	protoc  sso.proto --proto_path=protos/proto/sso/ --go_out=./protos/gen/go/sso --go_opt=paths=source_relative --go-grpc_out=./protos/gen/go/sso --go-grpc_opt=paths=source_relative