package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/handler"
	"github.com/palp1tate/brevinect/api/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/register", handler.Register)
		UserRouter.POST("/login", handler.UserLogin)
		UserRouter.GET("/get_user", middleware.JWTAuth(), handler.GetUser)
		UserRouter.PUT("/reset_password", middleware.JWTAuth(), handler.ResetPassword)
		UserRouter.PUT("/update_user", middleware.JWTAuth(), handler.UpdateUser)

		UserRouter.GET("get_all_company", handler.GetAllCompany)
	}
}
