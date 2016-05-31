
gogen:
	protoc --gogo_out=. --proto_path=${GOPATH}/src:. proto/*.proto

gen:
	protoc --go_out=plugins=grpc:. --proto_path=${GOPATH}/src:. proto/*.proto


deps:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/satori/go.uuid
	go get -u google.golang.org/grpc
	go get -u gopkg.in/urfave/cli.v2
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/gogoproto
