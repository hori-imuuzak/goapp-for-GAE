package main

import (
	"fmt"
	"net"
	"youtube-manager-go/api/usecase"
	pb "youtube-manager-go/gen/proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	port = 9000
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		glog.Fatalln(err)
		return
	}

	server := grpc.NewServer()

	testUsecase := usecase.NewTestUsecase()

	pb.RegisterTestServiceServer(server, testUsecase)

	server.Serve(lis)
}
