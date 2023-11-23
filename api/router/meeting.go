package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/handler"
	"github.com/palp1tate/brevinect/api/middleware"
)

func InitMeetingRouter(Router *gin.RouterGroup) {
	MeetingRouter := Router.Group("/meeting", middleware.JWTAuth())
	{
		MeetingRouter.GET("/get_room", handler.GetRoomByUser)
		MeetingRouter.GET("/get_room_list", handler.GetRoomListByUser)
	}
}
