build:
	go install && \
	protoc -I. -I`pwd`/proto \
	-I$(GOPATH)/src \
	-I/usr/local/protobuf/include \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--gator_out=golang:output schema.proto
