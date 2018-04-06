package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/jooita/GrpcRestExamples/echopb"
)

// join the two constants for convenience
var serveAddress = fmt.Sprintf("%v:%d", "localhost", 9090)

type server struct{}

// implements echo function of EchoServiceServer
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Echo: Received '%v'", in.Value)
	return &pb.StringMessage{Value: in.Value}, nil
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	var msg pb.StringMessage
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(msg.Value)
	w.Write([]byte(msg.Value))
}

func makeGRPCServer() *grpc.Server {

	//setup grpc server
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s
}

func serveGRPC(l net.Listener) {

	s := makeGRPCServer()

	if err := s.Serve(l); err != nil {
		log.Fatalf("While serving gRpc request: %v", err)
	}
}

func serveHTTP(l net.Listener) {
	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("While serving http request: %v", err)
	}
}

func tlsListener(tcpl net.Listener) (tlsl net.Listener) {

	serverCrt, err := ioutil.ReadFile("certs/server.crt")
	if err != nil {
		log.Fatal(err)
	}
	serverKey, err := ioutil.ReadFile("certs/server.key")
	if err != nil {
		log.Fatal(err)
	}

	pair, err := tls.X509KeyPair(serverCrt, serverKey)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{pair},
		NextProtos:   []string{"h2"},
	}

	tlsl = tls.NewListener(tcpl, config)
	return

}

func main() {

	// Create a listener at the desired port.
	tcpl, err := net.Listen("tcp", serveAddress)
	defer tcpl.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux object.
	tcpm := cmux.New(tcpl)

	// Declare the match for different services required.
	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcNotlsL := tcpm.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpNotlsL := tcpm.Match(cmux.HTTP1Fast())
	otherwiseL := tcpm.Match(cmux.Any())

	tlsl := tlsListener(otherwiseL)
	tlsm := cmux.New(tlsl)
	grpcTlsL := tlsm.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpTlsL := tlsm.Match(cmux.Any())

	// Link the endpoint to the handler function.
	// http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/v1/echo", EchoHandler)

	// Initialize the servers by passing in the custom listeners (sub-listeners).
	go serveGRPC(grpcNotlsL)
	go serveGRPC(grpcTlsL)

	go serveHTTP(httpNotlsL)
	go serveHTTP(httpTlsL)

	log.Println("grpc server started.")
	log.Println("http server started.")
	log.Println("Server listening on ", serveAddress)

	// Start cmux serving.
	go tcpm.Serve()
	if err := tlsm.Serve(); !strings.Contains(err.Error(),
		"use of closed network connection") {
		log.Fatal(err)
	}
}
