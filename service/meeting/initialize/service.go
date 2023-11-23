package initialize

import (
	"github.com/palp1tate/brevinect/proto/meeting"
	"github.com/palp1tate/brevinect/service/meeting/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	meetingProto.RegisterMeetingServiceServer(server, &handler.MeetingServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
