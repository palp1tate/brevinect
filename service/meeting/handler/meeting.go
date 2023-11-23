package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
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

func (s *MeetingServer) BookRoom(ctx context.Context, req *meetingProto.BookRoomRequest) (*empty.Empty, error) {
	book := model.Book{
		UserId:    int(req.UserId),
		RoomId:    int(req.RoomId),
		StartTime: req.Period.StartTime,
		EndTime:   req.Period.EndTime,
		Theme:     req.Theme,
	}
	err := dao.CreateBook(&book)
	if err != nil {
		return nil, status.Error(codes.Internal, "预定会议室失败")
	}
	return &empty.Empty{}, nil
}

func (s *MeetingServer) GetBookList(ctx context.Context, req *meetingProto.GetBookListRequest) (*meetingProto.GetBookListResponse, error) {
	books, pages, totalCount, err := dao.FindBookList(req.Page, req.PageSize, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "获取预定列表失败")
	}
	bookList := make([]*meetingProto.Book, len(books))
	for i, book := range books {
		bookList[i] = BookModelToResponse(book)
	}
	res := &meetingProto.GetBookListResponse{
		BookList:   bookList,
		Pages:      pages,
		TotalCount: totalCount,
	}
	return res, nil
}

func (s *MeetingServer) GetBook(ctx context.Context, req *meetingProto.GetBookRequest) (*meetingProto.GetBookResponse, error) {
	book, err := dao.FindBookById(int(req.BookId), int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该预定记录不存在")
	}
	res := &meetingProto.GetBookResponse{
		Book: BookModelToResponse(book),
	}
	return res, nil
}

func (s *MeetingServer) CancelBook(ctx context.Context, req *meetingProto.CancelBookRequest) (*empty.Empty, error) {
	book, err := dao.FindBookById(int(req.BookId), int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该预定记录不存在")
	}
	err = dao.DeleteBook(book)
	if err != nil {
		return nil, status.Error(codes.Internal, "取消预定失败")
	}
	return &empty.Empty{}, nil
}

func (s *MeetingServer) UpdateBook(ctx context.Context, req *meetingProto.UpdateBookRequest) (*empty.Empty, error) {
	book, err := dao.FindBookById(int(req.Book.Id), int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该预定记录不存在")
	}
	book.RoomId = int(req.Book.RoomId)
	book.StartTime = req.Book.Period.StartTime
	book.EndTime = req.Book.Period.EndTime
	book.Theme = req.Book.Theme
	err = dao.UpdateBook(book)
	if err != nil {
		return nil, status.Error(codes.Internal, "更新预定失败")
	}
	return &empty.Empty{}, nil
}

func RoomModelToResponse(room model.Room) *meetingProto.Room {
	photos, _ := dao.FindRoomPhoto(int(room.ID))
	books, _ := dao.FindBook(int(room.ID))
	bookedTime := make([]*meetingProto.Period, len(books))
	photoList := make([]string, len(photos))
	for i, photo := range photos {
		photoList[i] = photo.Url
	}
	for i, book := range books {
		bookedTime[i] = &meetingProto.Period{StartTime: book.StartTime, EndTime: book.EndTime}
	}
	return &meetingProto.Room{
		Id:         int64(room.ID),
		Name:       room.Name,
		Capacity:   int64(room.Capacity),
		Facility:   room.Facility,
		Location:   room.Location,
		Photo:      photoList,
		BookedTime: bookedTime,
	}
}

func BookModelToResponse(book model.Book) *meetingProto.Book {
	return &meetingProto.Book{
		Id:     int64(book.ID),
		RoomId: int64(book.RoomId),
		Period: &meetingProto.Period{StartTime: book.StartTime, EndTime: book.EndTime},
		Theme:  book.Theme,
	}
}
