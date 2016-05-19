
gen:
	protoc --go_out=plugins=grpc:. --proto_path=${GOPATH}/src:. proto/*.proto
