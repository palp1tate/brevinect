package dao

import (
	"github.com/palp1tate/brevinect/consts"
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/service/user/global"
)

func FindUserByMobile(mobile string) (user model.User, err error) {
	err = global.DB.Where("mobile = ? and role = ?", mobile, consts.User).First(&user).Error
	return
}

func FindUserById(id int) (user model.User, err error) {
	err = global.DB.Where("id = ? and role = ?", id, consts.User).First(&user).Error
	return
}

func CreateUser(user *model.User) (err error) {
	err = global.DB.Create(&user).Error
	return
}

func UpdatePassword(user model.User, password string) (err error) {
	err = global.DB.Model(&user).Update("password", password).Error
	return
}

func UpdateUser(user model.User) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func FindAllCompany() (companies []model.Company, err error) {
	err = global.DB.Order("id desc").Find(&companies).Error
	return
}
