package main

import (
	"github.com/jooita/GrpcRestExamples/clients/go/grpc"
	"github.com/jooita/GrpcRestExamples/clients/go/http"
)

func main() {
	grpc.EchoInsecure("localhost:9090")

	http.Echo("http://localhost:8080", http.StringMessage{Value: "hola, mundo!"})
}
