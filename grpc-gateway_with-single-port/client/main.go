package main

import (
	"github.com/jooita/GrpcRestExamples/clients/go/grpc"
	"github.com/jooita/GrpcRestExamples/clients/go/http"
	"io/ioutil"
	"log"
)

func main() {

	serverCrt, err := ioutil.ReadFile("../certs/server.crt")
	if err != nil {
		log.Fatal(err)
	}
	grpc.Echo("localhost:9090", serverCrt)

	http.Echo("https://localhost:9090", http.StringMessage{Value: "hola, mundo!"})

}
