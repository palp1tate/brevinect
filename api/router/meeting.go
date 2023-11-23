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

		MeetingRouter.POST("/book_room", handler.BookRoomByUser)
		MeetingRouter.GET("/get_book_list", handler.GetBookListByUser)
		MeetingRouter.GET("/get_book", handler.GetBookByUser)
		MeetingRouter.DELETE("/cancel_book", handler.CancelBookByUser)
		MeetingRouter.PUT("/update_book", handler.UpdateBookByUser)
	}
}
