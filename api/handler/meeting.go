package handler

import (
	"context"
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
