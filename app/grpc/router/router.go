package router

import (
	"google.golang.org/grpc"

	"github.com/agclqq/prow-pipeline/app/grpc/controller"
	"github.com/agclqq/prow-pipeline/app/grpc/pb/demo"
)

func Register(s *grpc.Server) {
	demo.RegisterDemoServer(s, &controller.Demo{Server: s})
}
