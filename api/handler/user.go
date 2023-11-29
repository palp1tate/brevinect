package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/brevinect/api/form"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/middleware"
	"github.com/palp1tate/brevinect/consts"
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/proto/user"
)

func Register(c *gin.Context) {
	registerForm := form.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	_, err := global.ThirdPartyServiceClient.CheckSmsCaptcha(context.Background(), &thirdProto.CheckSmsCaptchaRequest{
		Mobile:  registerForm.Mobile,
		Captcha: registerForm.Code,
		Type:    int64(consts.Register),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	res, err := global.UserServiceClient.Register(context.Background(), &userProto.RegisterRequest{
		Username: global.Filter.Replace(registerForm.Username, '*'),
		Mobile:   registerForm.Mobile,
		Password: registerForm.Password,
		Company:  int64(registerForm.Company),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	claims := middleware.CustomClaims{
		ID:      int(res.Id),
		Company: int(res.Company),
		Role:    consts.User,
	}
	j := middleware.NewJWT()
	token, err := j.CreateToken(claims)
	if err != nil {
		HandleHttpResponse(c, http.StatusInternalServerError, "生成token失败", nil, nil)
		return
	}
	HandleHttpResponse(c, http.StatusOK, "注册成功", token, nil)
	return
}

func UserLogin(c *gin.Context) {
	loginForm := form.LoginForm{}
	if err := c.ShouldBind(&loginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	var token string
	switch loginForm.Type {
	case consts.LoginByPassword:
		_, err := global.ThirdPartyServiceClient.CheckPicCaptcha(context.Background(), &thirdProto.CheckPicCaptchaRequest{
			CaptchaId: loginForm.CaptchaId,
			Captcha:   loginForm.Captcha,
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		res, err := global.UserServiceClient.LoginByPassword(context.Background(), &userProto.LoginByPasswordRequest{
			Mobile:   loginForm.Mobile,
			Password: loginForm.Password,
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		claims := middleware.CustomClaims{
			ID:      int(res.Id),
			Company: int(res.Company),
			Role:    consts.User,
		}
		j := middleware.NewJWT()
		token, err = j.CreateToken(claims)
		if err != nil {
			HandleHttpResponse(c, http.StatusInternalServerError, "生成token失败", nil, nil)
			return
		}
	case consts.LoginByCaptcha:
		_, err := global.ThirdPartyServiceClient.CheckSmsCaptcha(context.Background(), &thirdProto.CheckSmsCaptchaRequest{
			Mobile:  loginForm.Mobile,
			Captcha: loginForm.Captcha,
			Type:    int64(consts.Login),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		res, err := global.UserServiceClient.LoginBySMS(context.Background(), &userProto.LoginBySMSRequest{
			Mobile: loginForm.Mobile,
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		claims := middleware.CustomClaims{
			ID:      int(res.Id),
			Company: int(res.Company),
			Role:    consts.User,
		}
		j := middleware.NewJWT()
		token, err = j.CreateToken(claims)
		if err != nil {
			HandleHttpResponse(c, http.StatusInternalServerError, "生成token失败", nil, nil)
			return
		}
	}
	HandleHttpResponse(c, http.StatusOK, "登录成功", token, nil)
	return
}

func GetUserByUser(c *gin.Context) {
	userId := c.GetInt("id")
	res, err := global.UserServiceClient.GetUser(context.Background(), &userProto.GetUserRequest{
		Id: int64(userId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取用户信息成功", refreshedToken, res)
	return
}

func ResetPassword(c *gin.Context) {
	resetPasswordForm := form.ResetPasswordForm{}
	if err := c.ShouldBind(&resetPasswordForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	_, err := global.ThirdPartyServiceClient.CheckSmsCaptcha(context.Background(), &thirdProto.CheckSmsCaptchaRequest{
		Mobile:  resetPasswordForm.Mobile,
		Captcha: resetPasswordForm.Captcha,
		Type:    int64(consts.ResetPassword),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	_, err = global.UserServiceClient.ResetPassword(context.Background(), &userProto.ResetPasswordRequest{
		Mobile:   resetPasswordForm.Mobile,
		Password: resetPasswordForm.Password,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	HandleHttpResponse(c, http.StatusOK, "重置密码成功", nil, nil)
	return
}

func UpdateUserByUser(c *gin.Context) {
	updateUserForm := form.UpdateUserFormByUser{}
	if err := c.ShouldBind(&updateUserForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	userId := c.GetInt("id")
	_, err := global.UserServiceClient.UpdateUser(context.Background(), &userProto.UpdateUserRequest{
		Id:       int64(userId),
		Username: global.Filter.Replace(updateUserForm.Username, '*'),
		Avatar:   updateUserForm.Avatar,
		Face:     updateUserForm.Face,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "更新用户信息成功", refreshedToken, nil)
	return
}

func GetAllCompany(c *gin.Context) {
	res, err := global.UserServiceClient.GetAllCompany(context.Background(), &empty.Empty{})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	HandleHttpResponse(c, http.StatusOK, "获取所有公司成功", nil, res.Companies)
	return
}

func GetCompanyByUser(c *gin.Context) {
	cid := c.GetInt("company")
	res, err := global.UserServiceClient.GetCompany(context.Background(), &userProto.GetCompanyRequest{
		Id: int64(cid),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取公司信息成功", refreshedToken, res)
}

func UploadFace(c *gin.Context) {
	urlForm := form.UrlForm{}
	if err := c.ShouldBind(&urlForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	userId := c.GetInt("id")
	_, err := global.UserServiceClient.UploadFace(context.Background(), &userProto.UploadFaceRequest{
		Id:  int64(userId),
		Url: urlForm.Url,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "上传人脸成功", refreshedToken, nil)
	return
}
