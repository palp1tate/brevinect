package initialize

import (
	"io"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/service/third/global"
	"github.com/palp1tate/brevinect/service/third/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() (*grpc.Server, io.Closer) {
	tracer, closer := NewJaegerTracer(global.ServerConfig.Service.Name)
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	thirdProto.RegisterThirdPartyServiceServer(server, &handler.ThirdPartyServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server, closer
}
