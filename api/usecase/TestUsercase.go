package usecase

import (
	"context"
	pb "youtube-manager-go/gen/pb"
)

type testUsecase struct{}

func NewTestUsecase() *testUsecase {
	return &testUsecase{}
}

func (u *testUsecase) Test(c context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Message: "Test",
	}, nil
}
