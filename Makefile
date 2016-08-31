build:
	@mkdir bin
	@go build -o bin/tofu ./cmd/tofu

deps:
	go get -u github.com/satori/go.uuid
	go get -u google.golang.org/grpc
	go get -u gopkg.in/urfave/cli.v2
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/jsonpb
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/gogoproto

gen:
	protoc --gogo_out=plugins=grpc:. --proto_path=${GOPATH}/src:. proto/*.proto
