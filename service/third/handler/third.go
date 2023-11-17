package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/service/third/global"
	"github.com/palp1tate/brevinect/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ThirdPartyServer struct {
	thirdProto.UnimplementedThirdPartyServiceServer
}

func (s *ThirdPartyServer) GetPicCaptcha(ctx context.Context, req *empty.Empty) (*thirdProto.GetPicCaptchaResponse, error) {
	id, b64s, err := util.GeneratePicCaptcha()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "生成图片验证码失败:%s", err.Error())
	}
	res := &thirdProto.GetPicCaptchaResponse{
		CaptchaId: id,
		PicPath:   b64s,
	}
	return res, nil
}

func (s *ThirdPartyServer) CheckPicCaptcha(ctx context.Context, req *thirdProto.CheckPicCaptchaRequest) (*empty.Empty, error) {
	if ok := util.Store.Verify(req.CaptchaId, req.Captcha, true); !ok {
		return nil, status.Error(codes.Unauthenticated, "图片验证码错误")
	}
	return &empty.Empty{}, nil
}

func (s *ThirdPartyServer) GetSmsCaptcha(ctx context.Context, req *thirdProto.GetSmsCaptchaRequest) (*empty.Empty, error) {
	smsCode, err := util.SendSms(req.Mobile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "发送短信失败:%s", err.Error())
	}
	global.RedisClient.Set(context.Background(), fmt.Sprintf("%d-%s", req.Type, req.Mobile), smsCode,
		time.Duration(global.ServerConfig.Redis.Expiration)*time.Minute)
	return &empty.Empty{}, nil
}

func (s *ThirdPartyServer) CheckSmsCaptcha(ctx context.Context, req *thirdProto.CheckSmsCaptchaRequest) (*empty.Empty, error) {
	pattern := fmt.Sprintf("%d-%s", req.Type, req.Mobile)
	smsCode, err := global.RedisClient.Get(context.Background(), pattern).Result()
	if err == redis.Nil {
		return nil, status.Error(codes.Unauthenticated, "短信验证码已过期")
	} else {
		if smsCode != req.Captcha {
			return nil, status.Error(codes.Unauthenticated, "短信验证码错误")
		}
		global.RedisClient.Del(context.Background(), pattern)
	}
	return &empty.Empty{}, nil
}

func (s *ThirdPartyServer) UploadFile(ctx context.Context, req *thirdProto.UploadFileRequest) (*thirdProto.UploadFileResponse, error) {
	url, err := util.Upload(req.Data, req.Size, req.Suffix)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "上传文件失败:%s", err.Error())
	}
	res := &thirdProto.UploadFileResponse{
		Url: url,
	}
	return res, nil
}

func (s *ThirdPartyServer) DeleteFile(ctx context.Context, req *thirdProto.DeleteFileRequest) (*empty.Empty, error) {
	err := util.Delete(req.Url)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除文件失败:%s", err.Error())
	}
	return &empty.Empty{}, nil
}
