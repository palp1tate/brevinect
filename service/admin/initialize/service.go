package initialize

import (
	"github.com/palp1tate/brevinect/proto/admin"
	"github.com/palp1tate/brevinect/service/admin/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	adminProto.RegisterAdminServiceServer(server, &handler.AdminServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
