package person_service

import (
	"context"

	desc "github.com/ozonva/ova-person-api/pkg"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *PersonApi) CreatePersonV1(ctx context.Context, req *desc.CreatePersonRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a *PersonApi) DescribePersonV1(ctx context.Context, req *desc.DescribePersonRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a *PersonApi) ListPersonsPersonV1(ctx context.Context, req *desc.ListPersonsRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a *PersonApi) RemovePersonV1(ctx context.Context, req *desc.RemovePersonRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
