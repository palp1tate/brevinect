package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"path/filepath"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/brevinect/api/form"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/middleware"
	"github.com/palp1tate/brevinect/consts"
	"github.com/palp1tate/brevinect/proto/admin"
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/proto/user"
	"go.uber.org/zap"
)

func GetPicCaptcha(ctx *gin.Context) {
	res, err := global.ThirdPartyServiceClient.GetPicCaptcha(ctx, &empty.Empty{})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	HandleHttpResponse(ctx, http.StatusOK, "获取图片验证码成功", nil, res)
	return
}

func SendSms(c *gin.Context) {
	sendSmsForm := form.SendSmsForm{}
	if err := c.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	e, b := sentinel.Entry(consts.SMSResource, sentinel.WithTrafficType(base.Inbound), sentinel.WithArgs(sendSmsForm.Mobile))
	if b != nil {
		HandleHttpResponse(c, http.StatusTooManyRequests, "请求过于频繁，请稍后再试", nil, nil)
		return
	}
	defer e.Exit()
	if sendSmsForm.Role == consts.User {
		switch sendSmsForm.Type {
		case consts.Register:
			if res, err := global.UserServiceClient.CheckMobile(context.Background(), &userProto.CheckMobileRequest{
				Mobile: sendSmsForm.Mobile,
			}); err != nil {
				zap.S().Info("服务器内部错误")
				HandleGrpcErrorToHttp(c, err)
				return
			} else if res.Exist {
				zap.S().Info("该手机号已经被注册")
				HandleHttpResponse(c, http.StatusBadRequest, "该手机号已经被注册", nil, nil)
				return
			}
			_, err := global.ThirdPartyServiceClient.GetSmsCaptcha(context.Background(), &thirdProto.GetSmsCaptchaRequest{
				Mobile: sendSmsForm.Mobile,
				Type:   int64(consts.Register),
			})
			if err != nil {
				HandleGrpcErrorToHttp(c, err)
				return
			}
		case consts.Login:
			if res, err := global.UserServiceClient.CheckMobile(context.Background(), &userProto.CheckMobileRequest{
				Mobile: sendSmsForm.Mobile,
			}); err != nil {
				zap.S().Info("服务器内部错误")
				HandleGrpcErrorToHttp(c, err)
				return
			} else if !res.Exist {
				zap.S().Info("该手机号未注册")
				HandleHttpResponse(c, http.StatusBadRequest, "该手机号未注册", nil, nil)
				return
			}
			_, err := global.ThirdPartyServiceClient.GetSmsCaptcha(context.Background(), &thirdProto.GetSmsCaptchaRequest{
				Mobile: sendSmsForm.Mobile,
				Type:   int64(consts.Login),
			})
			if err != nil {
				HandleGrpcErrorToHttp(c, err)
				return
			}
		case consts.ResetPassword:
			if res, err := global.UserServiceClient.CheckMobile(context.Background(), &userProto.CheckMobileRequest{
				Mobile: sendSmsForm.Mobile,
			}); err != nil {
				zap.S().Info("服务器内部错误")
				HandleGrpcErrorToHttp(c, err)
				return
			} else if !res.Exist {
				zap.S().Info("该手机号未注册")
				HandleHttpResponse(c, http.StatusBadRequest, "该手机号未注册", nil, nil)
				return
			}
			_, err := global.ThirdPartyServiceClient.GetSmsCaptcha(context.Background(), &thirdProto.GetSmsCaptchaRequest{
				Mobile: sendSmsForm.Mobile,
				Type:   int64(consts.ResetPassword),
			})
			if err != nil {
				HandleGrpcErrorToHttp(c, err)
				return
			}
		}
	} else {
		if sendSmsForm.Type == consts.Login {
			if res, err := global.AdminServiceClient.CheckMobile(context.Background(), &adminProto.CheckMobileRequest{
				Mobile: sendSmsForm.Mobile,
			}); err != nil {
				zap.S().Info("服务器内部错误")
				HandleGrpcErrorToHttp(c, err)
				return
			} else if !res.Exist {
				zap.S().Info("非管理员账户")
				HandleHttpResponse(c, http.StatusBadRequest, "非管理员账户", nil, nil)
				return
			}
			_, err := global.ThirdPartyServiceClient.GetSmsCaptcha(context.Background(), &thirdProto.GetSmsCaptchaRequest{
				Mobile: sendSmsForm.Mobile,
				Type:   int64(consts.Login),
			})
			if err != nil {
				HandleGrpcErrorToHttp(c, err)
				return
			}
		}
	}
	HandleHttpResponse(c, http.StatusOK, "发送短信验证码成功", nil, nil)
	return
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		zap.S().Info("文件上传失败:", err)
		HandleHttpResponse(c, http.StatusBadRequest, "文件上传失败", nil, nil)
		return
	}
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)
	if err != nil {
		zap.S().Info("文件读取失败:", err)
		HandleHttpResponse(c, http.StatusBadRequest, "文件读取失败", nil, nil)
		return
	}
	res, err := global.ThirdPartyServiceClient.UploadFile(context.Background(), &thirdProto.UploadFileRequest{
		Data:   buffer.Bytes(),
		Size:   header.Size,
		Suffix: filepath.Ext(header.Filename),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "文件上传成功", refreshedToken, res.Url)
	return
}

func DeleteFile(c *gin.Context) {
	urlForm := form.UrlForm{}
	if err := c.ShouldBind(&urlForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	_, err := global.ThirdPartyServiceClient.DeleteFile(context.Background(), &thirdProto.DeleteFileRequest{
		Url: urlForm.Url,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "文件删除成功", refreshedToken, nil)
	return
}
