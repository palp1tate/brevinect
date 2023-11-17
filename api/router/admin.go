package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/handler"
	"github.com/palp1tate/brevinect/api/middleware"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	AdminRouter := Router.Group("/admin")
	{
		AdminRouter.POST("/login", handler.AdminLogin)
		AdminRouter.GET("/get_admin", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetAdmin)

		AdminRouter.GET("/get_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetCompany)
		AdminRouter.GET("/get_company_list", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetCompanyList)
		AdminRouter.POST("/add_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.AddCompany)
		AdminRouter.PUT("/update_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.UpdateCompany)
		AdminRouter.DELETE("/delete_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.DeleteCompany)

	}
}
