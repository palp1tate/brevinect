package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/handler"
	"github.com/palp1tate/brevinect/api/middleware"
)

func InitThirdPartyRouter(Router *gin.RouterGroup) {
	ThirdPartyRouter := Router.Group("/third_party")
	{
		ThirdPartyRouter.GET("/captcha", handler.GetPicCaptcha)
		ThirdPartyRouter.POST("/send_sms", handler.SendSms)

		ThirdPartyRouter.POST("upload_file", middleware.JWTAuth(), handler.UploadFile)
		ThirdPartyRouter.DELETE("delete_file", middleware.JWTAuth(), handler.DeleteFile)
	}
}
