package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/jooita/GrpcRestExamples/echopb"
)

type echoServer struct{}

func newEchoServer() pb.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func Run() error {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, newEchoServer())

	s.Serve(l)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
