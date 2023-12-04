package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/proto/meeting"
	"github.com/palp1tate/brevinect/service/meeting/dao"
	"github.com/xuri/excelize/v2"
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

func (s *MeetingServer) GetBookExcel(ctx context.Context, req *meetingProto.GetBookExcelRequest) (*meetingProto.GetBookExcelResponse, error) {
	// 获取用户的预定记录
	books, err := dao.FindBooks(int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, "获取预定记录失败")
	}
	// 获取用户信息
	user, err := dao.FindUserById(int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, "获取用户信息失败")
	}
	f := excelize.NewFile()
	// 创建一个新的Sheet
	sheet := "Sheet1"
	f.NewSheet(sheet)
	// 设置表头
	headers := []string{"会议室名称", "地点", "开始时间", "结束时间", "会议主题", "预定时间"}
	for i, h := range headers {
		colLetter := string(rune('A' + i)) // Excel列的字母标识
		cellName := colLetter + "1"        // 单元格名称，如"A1"
		f.SetCellValue(sheet, cellName, h)
		f.SetColWidth(sheet, colLetter, colLetter, 20)
	}

	// 将预定记录添加到Excel文件中
	for i, book := range books {
		row := i + 2
		// 获取会议室信息
		room, err := dao.FindRoomById(book.RoomId)
		if err != nil {
			return nil, status.Error(codes.Internal, "获取会议室信息失败")
		}
		values := []string{
			room.Name,
			room.Location,
			time.Unix(book.StartTime, 0).Format("2006-01-02 15:04"),
			time.Unix(book.EndTime, 0).Format("2006-01-02 15:04"),
			book.Theme,
			book.CreatedAt.Format("2006-01-02 15:04"),
		}
		for j, v := range values {
			colLetter := string(rune('A' + j))
			cellName := colLetter + strconv.Itoa(row)
			f.SetCellValue(sheet, cellName, v)
		}
	}

	// 将Excel文件转换为字节流
	excel, err := f.WriteToBuffer()
	if err != nil {
		return nil, status.Error(codes.Internal, "生成Excel文件失败")
	}

	// 创建响应
	res := &meetingProto.GetBookExcelResponse{
		Excel:    excel.Bytes(),
		FileName: fmt.Sprintf("%s的会议室预定记录表.xlsx", user.Username),
		Size:     int64(excel.Len()),
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
		Id:       int64(book.ID),
		RoomId:   int64(book.RoomId),
		Period:   &meetingProto.Period{StartTime: book.StartTime, EndTime: book.EndTime},
		Theme:    book.Theme,
		BookTime: book.CreatedAt.Unix(),
	}
}
