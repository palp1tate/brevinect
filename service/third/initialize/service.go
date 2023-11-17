package initialize

import (
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/service/third/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	thirdProto.RegisterThirdPartyServiceServer(server, &handler.ThirdPartyServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
