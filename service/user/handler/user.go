package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/proto/user"
	"github.com/palp1tate/brevinect/service/user/dao"
	"github.com/palp1tate/go-crypto-guard/pbkdf2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
}

func (s *UserServer) CheckMobile(ctx context.Context, req *userProto.CheckMobileRequest) (*userProto.CheckMobileResponse, error) {
	_, err := dao.FindUserByMobile(req.Mobile)
	res := &userProto.CheckMobileResponse{
		Exist: false,
	}
	if err == nil {
		res.Exist = true
	}
	return res, nil
}

func (s *UserServer) Register(ctx context.Context, req *userProto.RegisterRequest) (*userProto.RegisterResponse, error) {
	password, _ := pwd.GenSHA512(req.Password, 16, 32, 50)
	user := model.User{
		Mobile:    req.Mobile,
		Password:  password,
		Username:  req.Username,
		CompanyId: int(req.Company),
	}
	err := dao.CreateUser(&user)
	if err != nil {
		return nil, status.Error(codes.Internal, "用户注册失败")
	}
	res := &userProto.RegisterResponse{
		Id:      int64(user.ID),
		Company: int64(user.CompanyId),
	}
	return res, nil
}

func (s *UserServer) LoginByPassword(ctx context.Context, req *userProto.LoginByPasswordRequest) (*userProto.LoginResponse, error) {
	user, err := dao.FindUserByMobile(req.Mobile)
	if err != nil {
		return nil, status.Error(codes.NotFound, "该手机号未注册")
	}
	if ok, _ := pwd.VerifySHA512(req.Password, user.Password); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "密码错误")
	}
	res := &userProto.LoginResponse{
		Id:      int64(user.ID),
		Company: int64(user.CompanyId),
	}
	return res, nil
}

func (s *UserServer) LoginBySMS(ctx context.Context, req *userProto.LoginBySMSRequest) (*userProto.LoginResponse, error) {
	user, err := dao.FindUserByMobile(req.Mobile)
	if err != nil {
		return nil, status.Error(codes.NotFound, "该手机号未注册")
	}
	res := &userProto.LoginResponse{
		Id:      int64(user.ID),
		Company: int64(user.CompanyId),
	}
	return res, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *userProto.GetUserRequest) (*userProto.GetUserResponse, error) {
	user, err := dao.FindUserById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该用户不存在")
	}
	res := &userProto.GetUserResponse{
		Id:       int64(user.ID),
		Username: user.Username,
		Mobile:   user.Mobile,
		Company:  int64(user.CompanyId),
		Avatar:   user.Avatar,
		Face:     user.Face,
	}
	return res, nil
}

func (s *UserServer) ResetPassword(ctx context.Context, req *userProto.ResetPasswordRequest) (*empty.Empty, error) {
	user, err := dao.FindUserByMobile(req.Mobile)
	if err != nil {
		return nil, status.Error(codes.NotFound, "该手机号未注册")
	}
	password, _ := pwd.GenSHA512(req.Password, 16, 32, 50)
	err = dao.UpdatePassword(user, password)
	if err != nil {
		return nil, status.Error(codes.Internal, "重置密码失败")
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *userProto.UpdateUserRequest) (*empty.Empty, error) {
	user, err := dao.FindUserById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该用户不存在")
	}
	user.Username = req.Username
	user.Avatar = req.Avatar
	user.Face = req.Face
	err = dao.UpdateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "更新用户信息失败")
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) GetAllCompany(ctx context.Context, req *empty.Empty) (*userProto.GetAllCompanyResponse, error) {
	companies, err := dao.FindAllCompany()
	if err != nil {
		return nil, status.Error(codes.Internal, "获取公司列表失败")
	}
	companyList := make([]*userProto.Company, len(companies))
	for i, company := range companies {
		companyList[i] = CompanyModelToResponse(company)
	}
	res := &userProto.GetAllCompanyResponse{
		Companies: companyList,
	}
	return res, nil
}

func (s *UserServer) GetCompany(ctx context.Context, req *userProto.GetCompanyRequest) (*userProto.GetCompanyResponse, error) {
	company, err := dao.FindCompanyById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该公司不存在")
	}
	res := &userProto.GetCompanyResponse{
		Id:             int64(company.ID),
		Name:           company.Name,
		Address:        company.Address,
		OfficialMobile: company.OfficialMobile,
		OfficialSite:   company.OfficialSite,
		CompanyType:    company.CompanyType,
		Introduction:   company.Introduction,
		Picture:        company.Picture,
	}
	return res, nil
}

func (s *UserServer) UploadFace(ctx context.Context, req *userProto.UploadFaceRequest) (*empty.Empty, error) {
	user, err := dao.FindUserById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该用户不存在")
	}
	user.Face = req.Url
	err = dao.UpdateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "上传人脸失败")
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) CheckFace(ctx context.Context, req *userProto.CheckFaceRequest) (*empty.Empty, error) {
	user, err := dao.FindUserById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该用户不存在")
	}
	if user.Face == "" {
		return nil, status.Error(codes.Internal, "该用户未上传人脸")
	}
	return &empty.Empty{}, nil
}

func CompanyModelToResponse(company model.Company) *userProto.Company {
	return &userProto.Company{
		Id:   int64(company.ID),
		Name: company.Name,
	}
}
