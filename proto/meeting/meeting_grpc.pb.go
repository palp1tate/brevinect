// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: meeting.proto

package meetingProto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeetingServiceClient is the client API for MeetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeetingServiceClient interface {
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
	GetRoomList(ctx context.Context, in *GetRoomListRequest, opts ...grpc.CallOption) (*GetRoomListResponse, error)
	BookRoom(ctx context.Context, in *BookRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelBook(ctx context.Context, in *CancelBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
	GetBookExcel(ctx context.Context, in *GetBookExcelRequest, opts ...grpc.CallOption) (*GetBookExcelResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type meetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingServiceClient(cc grpc.ClientConnInterface) MeetingServiceClient {
	return &meetingServiceClient{cc}
}

func (c *meetingServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetRoomList(ctx context.Context, in *GetRoomListRequest, opts ...grpc.CallOption) (*GetRoomListResponse, error) {
	out := new(GetRoomListResponse)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/GetRoomList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) BookRoom(ctx context.Context, in *BookRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/BookRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) CancelBook(ctx context.Context, in *CancelBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/CancelBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error) {
	out := new(GetBookListResponse)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) GetBookExcel(ctx context.Context, in *GetBookExcelRequest, opts ...grpc.CallOption) (*GetBookExcelResponse, error) {
	out := new(GetBookExcelResponse)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/GetBookExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/meetingProto.MeetingService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingServiceServer is the server API for MeetingService service.
// All implementations must embed UnimplementedMeetingServiceServer
// for forward compatibility
type MeetingServiceServer interface {
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
	GetRoomList(context.Context, *GetRoomListRequest) (*GetRoomListResponse, error)
	BookRoom(context.Context, *BookRoomRequest) (*emptypb.Empty, error)
	CancelBook(context.Context, *CancelBookRequest) (*emptypb.Empty, error)
	GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	GetBookExcel(context.Context, *GetBookExcelRequest) (*GetBookExcelResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMeetingServiceServer()
}

// UnimplementedMeetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingServiceServer struct {
}

func (UnimplementedMeetingServiceServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedMeetingServiceServer) GetRoomList(context.Context, *GetRoomListRequest) (*GetRoomListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (UnimplementedMeetingServiceServer) BookRoom(context.Context, *BookRoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookRoom not implemented")
}
func (UnimplementedMeetingServiceServer) CancelBook(context.Context, *CancelBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBook not implemented")
}
func (UnimplementedMeetingServiceServer) GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}
func (UnimplementedMeetingServiceServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedMeetingServiceServer) GetBookExcel(context.Context, *GetBookExcelRequest) (*GetBookExcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookExcel not implemented")
}
func (UnimplementedMeetingServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedMeetingServiceServer) mustEmbedUnimplementedMeetingServiceServer() {}

// UnsafeMeetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingServiceServer will
// result in compilation errors.
type UnsafeMeetingServiceServer interface {
	mustEmbedUnimplementedMeetingServiceServer()
}

func RegisterMeetingServiceServer(s grpc.ServiceRegistrar, srv MeetingServiceServer) {
	s.RegisterService(&MeetingService_ServiceDesc, srv)
}

func _MeetingService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/GetRoomList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetRoomList(ctx, req.(*GetRoomListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_BookRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).BookRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/BookRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).BookRoom(ctx, req.(*BookRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_CancelBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).CancelBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/CancelBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).CancelBook(ctx, req.(*CancelBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetBookList(ctx, req.(*GetBookListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_GetBookExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookExcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).GetBookExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/GetBookExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).GetBookExcel(ctx, req.(*GetBookExcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeetingService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meetingProto.MeetingService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeetingService_ServiceDesc is the grpc.ServiceDesc for MeetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meetingProto.MeetingService",
	HandlerType: (*MeetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoom",
			Handler:    _MeetingService_GetRoom_Handler,
		},
		{
			MethodName: "GetRoomList",
			Handler:    _MeetingService_GetRoomList_Handler,
		},
		{
			MethodName: "BookRoom",
			Handler:    _MeetingService_BookRoom_Handler,
		},
		{
			MethodName: "CancelBook",
			Handler:    _MeetingService_CancelBook_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _MeetingService_GetBookList_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _MeetingService_GetBook_Handler,
		},
		{
			MethodName: "GetBookExcel",
			Handler:    _MeetingService_GetBookExcel_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _MeetingService_UpdateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meeting.proto",
}
