package handler

import (
	"context"

	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/proto/meeting"
	"github.com/palp1tate/brevinect/service/meeting/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MeetingServer struct {
	meetingProto.UnimplementedMeetingServiceServer
}

func (s *MeetingServer) GetRoom(ctx context.Context, req *meetingProto.GetRoomRequest) (*meetingProto.GetRoomResponse, error) {
	room, err := dao.FindRoomById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该会议室不存在")
	}
	res := &meetingProto.GetRoomResponse{
		Room: RoomModelToResponse(room),
	}
	return res, nil
}

func (s *MeetingServer) GetRoomList(ctx context.Context, req *meetingProto.GetRoomListRequest) (*meetingProto.GetRoomListResponse, error) {
	rooms, pages, totalCount, err := dao.FindRoomList(req.Page, req.PageSize, req.Company)
	if err != nil {
		return nil, status.Error(codes.Internal, "获取会议室列表失败")
	}
	roomList := make([]*meetingProto.Room, len(rooms))
	for i, room := range rooms {
		roomList[i] = RoomModelToResponse(room)
	}
	res := &meetingProto.GetRoomListResponse{
		RoomList:   roomList,
		Pages:      pages,
		TotalCount: totalCount,
	}
	return res, nil
}

func RoomModelToResponse(room model.Room) *meetingProto.Room {
	photos, _ := dao.FindRoomPhoto(int(room.ID))
	photoList := make([]string, len(photos))
	for i, photo := range photos {
		photoList[i] = photo.Url
	}
	return &meetingProto.Room{
		Id:       int64(room.ID),
		Name:     room.Name,
		Capacity: int64(room.Capacity),
		Facility: room.Facility,
		Location: room.Location,
		Photo:    photoList,
	}
}
