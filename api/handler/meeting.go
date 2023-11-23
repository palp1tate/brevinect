package handler

import (
	"context"
	"github.com/palp1tate/brevinect/api/form"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/middleware"
	"github.com/palp1tate/brevinect/proto/meeting"
	"github.com/palp1tate/brevinect/util"
)

func GetRoomByUser(c *gin.Context) {
	roomId, _ := strconv.ParseInt(c.Query("rid"), 10, 64)
	if roomId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "rid不能为空", nil, nil)
		return
	}
	res, err := global.MeetingServiceClient.GetRoom(context.Background(), &meetingProto.GetRoomRequest{
		Id: roomId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取会议室信息成功", refreshedToken, res.Room)
	return
}

func GetRoomListByUser(c *gin.Context) {
	company := c.GetInt("company")
	page, pageSize := util.ParsePageAndPageSize(c.Query("page"), c.Query("pageSize"))
	res, err := global.MeetingServiceClient.GetRoomList(context.Background(), &meetingProto.GetRoomListRequest{
		Company:  int64(company),
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
	HandleHttpResponse(c, http.StatusOK, "获取会议室列表成功", refreshedToken, res)
	return
}

func BookRoomByUser(c *gin.Context) {
	bookRoomForm := form.BookRoomForm{}
	if err := c.ShouldBind(&bookRoomForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	userId := c.GetInt("id")
	_, err := global.MeetingServiceClient.BookRoom(context.Background(), &meetingProto.BookRoomRequest{
		UserId: int64(userId),
		RoomId: bookRoomForm.RoomId,
		Period: &meetingProto.Period{
			StartTime: bookRoomForm.StartTime,
			EndTime:   bookRoomForm.EndTime,
		},
		Theme: bookRoomForm.Theme,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "预定会议室成功", refreshedToken, nil)
	return
}

func GetBookListByUser(c *gin.Context) {
	page, pageSize := util.ParsePageAndPageSize(c.Query("page"), c.Query("pageSize"))
	userId := c.GetInt("id")
	res, err := global.MeetingServiceClient.GetBookList(context.Background(), &meetingProto.GetBookListRequest{
		UserId:   int64(userId),
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
	HandleHttpResponse(c, http.StatusOK, "获取预定记录列表成功", refreshedToken, res)
	return
}

func GetBookByUser(c *gin.Context) {
	bookId, _ := strconv.ParseInt(c.Query("bid"), 10, 64)
	if bookId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "bid不能为空", nil, nil)
		return
	}
	userId := c.GetInt("id")
	res, err := global.MeetingServiceClient.GetBook(context.Background(), &meetingProto.GetBookRequest{
		UserId: int64(userId),
		BookId: bookId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取预定记录成功", refreshedToken, res.Book)
	return
}

func CancelBookByUser(c *gin.Context) {
	bookId, _ := strconv.ParseInt(c.Query("bid"), 10, 64)
	if bookId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "bid不能为空", nil, nil)
		return
	}
	userId := c.GetInt("id")
	_, err := global.MeetingServiceClient.CancelBook(context.Background(), &meetingProto.CancelBookRequest{
		UserId: int64(userId),
		BookId: bookId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "取消预定成功", refreshedToken, nil)
	return
}

func UpdateBookByUser(c *gin.Context) {
	updateBookForm := form.UpdateBookForm{}
	if err := c.ShouldBind(&updateBookForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	userId := c.GetInt("id")
	_, err := global.MeetingServiceClient.UpdateBook(context.Background(), &meetingProto.UpdateBookRequest{
		UserId: int64(userId),
		Book: &meetingProto.Book{
			Id:     updateBookForm.BookId,
			RoomId: updateBookForm.RoomId,
			Theme:  updateBookForm.Theme,
			Period: &meetingProto.Period{StartTime: updateBookForm.StartTime, EndTime: updateBookForm.EndTime},
		},
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "更新预定成功", refreshedToken, nil)
	return
}
