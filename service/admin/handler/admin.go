package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/proto/admin"
	"github.com/palp1tate/brevinect/service/admin/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminServer struct {
	adminProto.UnimplementedAdminServiceServer
}

func (s *AdminServer) CheckMobile(ctx context.Context, req *adminProto.CheckMobileRequest) (*adminProto.CheckMobileResponse, error) {
	_, err := dao.FindAdminByMobile(req.Mobile)
	res := &adminProto.CheckMobileResponse{
		Exist: false,
	}
	if err == nil {
		res.Exist = true
	}
	return res, nil
}

func (s *AdminServer) LoginBySMS(ctx context.Context, req *adminProto.LoginBySMSRequest) (*adminProto.LoginResponse, error) {
	admin, err := dao.FindAdminByMobile(req.Mobile)
	if err != nil {
		return nil, status.Error(codes.NotFound, "非管理员账户")
	}
	res := &adminProto.LoginResponse{
		Id:      int64(admin.ID),
		Company: int64(admin.CompanyId),
	}
	return res, nil
}

func (s *AdminServer) LoginByPassword(ctx context.Context, req *adminProto.LoginByPasswordRequest) (*adminProto.LoginResponse, error) {
	admin, err := dao.FindAdminByMobile(req.Mobile)
	if err != nil {
		return nil, status.Error(codes.NotFound, "非管理员账户")
	}
	if admin.Password != req.Password {
		return nil, status.Errorf(codes.Unauthenticated, "密码错误")
	}
	res := &adminProto.LoginResponse{
		Id:      int64(admin.ID),
		Company: int64(admin.CompanyId),
	}
	return res, nil
}

func (s *AdminServer) GetAdmin(ctx context.Context, req *adminProto.GetAdminRequest) (*adminProto.GetAdminResponse, error) {
	admin, err := dao.FindAdminById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该管理员不存在")
	}
	res := &adminProto.GetAdminResponse{
		Id:       int64(admin.ID),
		Username: admin.Username,
		Mobile:   admin.Mobile,
		Company:  int64(admin.CompanyId),
		Avatar:   admin.Avatar,
	}
	return res, nil
}

func (s *AdminServer) AddCompany(ctx context.Context, req *adminProto.AddCompanyRequest) (*empty.Empty, error) {
	_, err := dao.FindCompanyByName(req.Name)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "该公司已存在")
	}
	company := model.Company{
		Name:           req.Name,
		Address:        req.Address,
		OfficialMobile: req.OfficialMobile,
		OfficialSite:   req.OfficialSite,
		CompanyType:    req.CompanyType,
		Introduction:   req.Introduction,
		Picture:        req.Picture,
	}
	err = dao.CreateCompany(&company)
	if err != nil {
		return nil, status.Error(codes.Internal, "新增公司失败")
	}
	return &empty.Empty{}, nil
}

func (s *AdminServer) UpdateCompany(ctx context.Context, req *adminProto.UpdateCompanyRequest) (*empty.Empty, error) {
	company, err := dao.FindCompanyByName(req.Company.Name)
	if err != nil {
		return nil, status.Error(codes.NotFound, "该公司不存在")
	}
	company.Name = req.Company.Name
	company.Address = req.Company.Address
	company.OfficialMobile = req.Company.OfficialMobile
	company.OfficialSite = req.Company.OfficialSite
	company.CompanyType = req.Company.CompanyType
	company.Introduction = req.Company.Introduction
	company.Picture = req.Company.Picture
	err = dao.UpdateCompany(company)
	if err != nil {
		return nil, status.Error(codes.Internal, "更新公司信息失败")
	}
	return &empty.Empty{}, nil
}

func (s *AdminServer) DeleteCompany(ctx context.Context, req *adminProto.DeleteCompanyRequest) (*empty.Empty, error) {
	company, err := dao.FindCompanyById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该公司不存在")
	}
	err = dao.DeleteCompany(company)
	if err != nil {
		return nil, status.Error(codes.Internal, "删除公司失败")
	}
	return &empty.Empty{}, nil
}

func (s *AdminServer) GetCompany(ctx context.Context, req *adminProto.GetCompanyRequest) (*adminProto.GetCompanyResponse, error) {
	company, err := dao.FindCompanyById(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "该公司不存在")
	}
	res := &adminProto.GetCompanyResponse{
		Company: CompanyModelToResponse(company),
	}
	return res, nil
}

func (s *AdminServer) GetCompanyList(ctx context.Context, req *adminProto.GetCompanyListRequest) (*adminProto.GetCompanyListResponse, error) {
	companies, pages, totalCount, err := dao.FindCompanyList(req.Page, req.PageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, "获取公司列表失败")
	}
	companyList := make([]*adminProto.Company, len(companies))
	for i, company := range companies {
		companyList[i] = CompanyModelToResponse(company)
	}
	res := &adminProto.GetCompanyListResponse{
		CompanyList: companyList,
		Pages:       pages,
		TotalCount:  totalCount,
	}
	return res, nil
}

func CompanyModelToResponse(company model.Company) *adminProto.Company {
	return &adminProto.Company{
		Id:             int64(company.ID),
		Name:           company.Name,
		Address:        company.Address,
		OfficialMobile: company.OfficialMobile,
		OfficialSite:   company.OfficialSite,
		CompanyType:    company.CompanyType,
		Introduction:   company.Introduction,
		Picture:        company.Picture,
	}
}
