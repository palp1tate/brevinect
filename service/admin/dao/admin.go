package dao

import (
	"github.com/palp1tate/brevinect/consts"
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/service/admin/global"
)

func FindAdminByMobile(mobile string) (user model.User, err error) {
	err = global.DB.Where("mobile = ? and role = ?", mobile, consts.Admin).First(&user).Error
	return
}

func FindAdminById(id int) (user model.User, err error) {
	err = global.DB.Where("id = ? and role = ?", id, consts.Admin).First(&user).Error
	return
}

func CreateCompany(company *model.Company) (err error) {
	err = global.DB.Create(&company).Error
	return
}

func FindCompanyById(id int) (company model.Company, err error) {
	err = global.DB.Where("id = ?", id).First(&company).Error
	return
}

func FindCompanyByName(name string) (company model.Company, err error) {
	err = global.DB.Where("name = ?", name).First(&company).Error
	return
}

func UpdateCompany(company model.Company) (err error) {
	err = global.DB.Save(&company).Error
	return
}

func DeleteCompany(company model.Company) (err error) {
	err = global.DB.Unscoped().Delete(&company).Error
	return
}

func FindCompanyList(page int64, pageSize int64) (companies []model.Company, pages int64, totalCount int64, err error) {
	err = global.DB.Model(&model.Company{}).Count(&totalCount).Order("id desc").
		Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&companies).Error
	pages = totalCount / pageSize
	if totalCount%(pageSize) != 0 {
		pages++
	}
	return
}
