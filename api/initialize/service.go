package initialize

import (
	"fmt"
	"sync"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/proto/admin"
	"github.com/palp1tate/brevinect/proto/meeting"
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServiceConn() {
	consul := global.ServerConfig.Consul
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		userConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",
				consul.Host, consul.Port, global.ServerConfig.Service.User),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		)
		if err != nil {
			zap.S().Fatal("连接用户服务失败")
		}
		userServiceClient := userProto.NewUserServiceClient(userConn)
		global.UserServiceClient = userServiceClient
	}()

	go func() {
		defer wg.Done()
		adminConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",
				consul.Host, consul.Port, global.ServerConfig.Service.Admin),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		)
		if err != nil {
			zap.S().Fatal("连接管理员服务失败")
		}
		adminServiceClient := adminProto.NewAdminServiceClient(adminConn)
		global.AdminServiceClient = adminServiceClient
	}()

	go func() {
		defer wg.Done()
		meetingConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",
				consul.Host, consul.Port, global.ServerConfig.Service.Meeting),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		)
		if err != nil {
			zap.S().Fatal("连接会议服务失败")
		}
		meetingServiceClient := meetingProto.NewMeetingServiceClient(meetingConn)
		global.MeetingServiceClient = meetingServiceClient
	}()

	go func() {
		defer wg.Done()
		thirdConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",
				consul.Host, consul.Port, global.ServerConfig.Service.ThirdParty),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
		if err != nil {
			zap.S().Fatal("连接第三方服务失败")
		}
		thirdPartyServiceClient := thirdProto.NewThirdPartyServiceClient(thirdConn)
		global.ThirdPartyServiceClient = thirdPartyServiceClient
	}()

	wg.Wait()
}
