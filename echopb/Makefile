PROTO = echo.proto
NAME = echo
all: go python gw swagger

install:
	#python-grpc
	pip install googleapis-common-protos
	#go-grpc
	go get google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go	
	#grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

go:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. \
	$(PROTO)
python:
	python -m grpc_tools.protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--python_out=. --grpc_python_out=. \
	$(PROTO)
gw:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	$(PROTO)

swagger:
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. \
	$(PROTO)

clean:
	@rm -rf $(NAME).pb.go
	@rm -rf $(NAME)_pb2.py
	@rm -rf $(NAME)_pb2.pyc
	@rm -rf $(NAME)_pb2_grpc.py
	@rm -rf $(NAME)_pb2_grpc.pyc
	@rm -rf $(NAME).pb.gw.go
