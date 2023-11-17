package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/palp1tate/brevinect/util"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/form"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/middleware"
	"github.com/palp1tate/brevinect/consts"
	"github.com/palp1tate/brevinect/proto/admin"
	"github.com/palp1tate/brevinect/proto/third"
)

func AdminLogin(c *gin.Context) {
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
		res, err := global.AdminServiceClient.LoginByPassword(context.Background(), &adminProto.LoginByPasswordRequest{
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
			Role:    consts.Admin,
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
		res, err := global.AdminServiceClient.LoginBySMS(context.Background(), &adminProto.LoginBySMSRequest{
			Mobile: loginForm.Mobile,
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		claims := middleware.CustomClaims{
			ID:      int(res.Id),
			Company: int(res.Company),
			Role:    consts.Admin,
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

func GetAdmin(c *gin.Context) {
	adminId := c.GetInt("id")
	res, err := global.AdminServiceClient.GetAdmin(context.Background(), &adminProto.GetAdminRequest{
		Id: int64(adminId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取管理员信息成功", refreshedToken, res)
	return
}

func GetCompany(c *gin.Context) {
	companyId, _ := strconv.ParseInt(c.Query("cid"), 10, 64)
	if companyId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "cid不能为空", nil, nil)
		return
	}
	res, err := global.AdminServiceClient.GetCompany(context.Background(), &adminProto.GetCompanyRequest{
		Id: companyId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取公司信息成功", refreshedToken, res.Company)
	return
}

func GetCompanyList(c *gin.Context) {
	page, pageSize := util.ParsePageAndPageSize(c.Query("page"), c.Query("pageSize"))
	res, err := global.AdminServiceClient.GetCompanyList(context.Background(), &adminProto.GetCompanyListRequest{
		Page:     int64(page),
		PageSize: int64(pageSize),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取公司列表成功", refreshedToken, res)
}

func AddCompany(c *gin.Context) {
	addCompanyForm := form.AddCompanyForm{}
	if err := c.ShouldBind(&addCompanyForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	_, err := global.AdminServiceClient.AddCompany(context.Background(), &adminProto.AddCompanyRequest{
		Name:           addCompanyForm.Name,
		Address:        addCompanyForm.Address,
		Introduction:   addCompanyForm.Introduction,
		OfficialMobile: addCompanyForm.OfficialMobile,
		OfficialSite:   addCompanyForm.OfficialSite,
		CompanyType:    addCompanyForm.CompanyType,
		Picture:        addCompanyForm.Picture,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "添加公司成功", refreshedToken, nil)
	return
}

func UpdateCompany(c *gin.Context) {
	updateCompanyForm := form.UpdateCompanyForm{}
	if err := c.ShouldBind(&updateCompanyForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	_, err := global.AdminServiceClient.UpdateCompany(context.Background(), &adminProto.UpdateCompanyRequest{
		Company: &adminProto.Company{
			Id:             int64(updateCompanyForm.Id),
			Name:           updateCompanyForm.Name,
			Address:        updateCompanyForm.Address,
			Introduction:   updateCompanyForm.Introduction,
			OfficialMobile: updateCompanyForm.OfficialMobile,
			OfficialSite:   updateCompanyForm.OfficialSite,
			CompanyType:    updateCompanyForm.CompanyType,
			Picture:        updateCompanyForm.Picture,
		},
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "更新公司信息成功", refreshedToken, nil)
	return
}

func DeleteCompany(c *gin.Context) {
	companyId, _ := strconv.ParseInt(c.Query("cid"), 10, 64)
	if companyId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "cid不能为空", nil, nil)
		return
	}
	_, err := global.AdminServiceClient.DeleteCompany(context.Background(), &adminProto.DeleteCompanyRequest{
		Id: companyId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "删除公司成功", refreshedToken, nil)
	return
}
