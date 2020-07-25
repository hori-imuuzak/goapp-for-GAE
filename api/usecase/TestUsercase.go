package usecase

import (
	"context"
	pb "youtube-manager-go/gen/proto"
)

type testUsecase struct{}

func NewTestUsecase() *testUsecase {
	return &testUsecase{}
}

func (u *testUsecase) Test(c context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Message: "Test Deployment!",
	}, nil
}

func (u *testUsecase) Hello(c context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, GCP!!",
	}, nil
}
