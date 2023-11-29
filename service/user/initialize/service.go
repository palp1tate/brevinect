package initialize

import (
	"io"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/palp1tate/brevinect/proto/user"
	"github.com/palp1tate/brevinect/service/user/global"
	"github.com/palp1tate/brevinect/service/user/handler"
	"google.golang.org/grpc"
)

func InitGRPC() (*grpc.Server, io.Closer) {
	tracer, closer := NewJaegerTracer(global.ServerConfig.Service.Name)
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	userProto.RegisterUserServiceServer(server, &handler.UserServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server, closer
}
