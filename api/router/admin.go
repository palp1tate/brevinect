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

		AdminRouter.GET("/get_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetCompanyByAdmin)
		AdminRouter.GET("/get_company_list", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetCompanyList)
		AdminRouter.POST("/add_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.AddCompany)
		AdminRouter.PUT("/update_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.UpdateCompany)
		AdminRouter.DELETE("/delete_company", middleware.JWTAuth(), middleware.AdminAuth(), handler.DeleteCompany)

		AdminRouter.GET("/get_room", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetRoomByAdmin)
		AdminRouter.GET("/get_room_list", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetRoomListByAdmin)
		AdminRouter.POST("/add_room", middleware.JWTAuth(), middleware.AdminAuth(), handler.AddRoom)
		AdminRouter.PUT("/update_room", middleware.JWTAuth(), middleware.AdminAuth(), handler.UpdateRoom)
		AdminRouter.DELETE("/delete_room", middleware.JWTAuth(), middleware.AdminAuth(), handler.DeleteRoom)

		AdminRouter.GET("/get_user", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetUserByAdmin)
		AdminRouter.GET("/get_user_list", middleware.JWTAuth(), middleware.AdminAuth(), handler.GetUserListByAdmin)
		AdminRouter.PUT("/update_user", middleware.JWTAuth(), middleware.AdminAuth(), handler.UpdateUserByAdmin)
		AdminRouter.DELETE("/delete_user", middleware.JWTAuth(), middleware.AdminAuth(), handler.DeleteUser)

	}
}
