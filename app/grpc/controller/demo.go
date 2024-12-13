package controller

import (
	"context"

	"google.golang.org/grpc"

	"github.com/agclqq/prow-pipeline/app/grpc/pb/demo"
)

type Demo struct {
	Server *grpc.Server
	demo.UnimplementedDemoServer
}

func (d *Demo) Foo(context.Context, *demo.DemoReq) (*demo.DemoRes, error) {
	return &demo.DemoRes{Name: "Hello World"}, nil
}
