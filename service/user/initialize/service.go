package initialize

import (
	"github.com/palp1tate/brevinect/proto/user"
	"github.com/palp1tate/brevinect/service/user/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	userProto.RegisterUserServiceServer(server, &handler.UserServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
