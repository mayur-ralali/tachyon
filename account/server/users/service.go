package users

import (
	"context"

	"github.com/mayur-ralali/tachyon/account/pb"
)

type Service struct{}

func (u *Service) Add(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{Result: r.GetA() + r.GetB()}, nil
}

func (u *Service) Multiply(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{Result: r.GetA() * r.GetB()}, nil
}
