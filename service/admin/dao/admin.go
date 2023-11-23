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

func CreateRoom(room *model.Room, photo []string) (err error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Create(&room).Error; err != nil {
		tx.Rollback()
		return
	}
	for _, p := range photo {
		photoModel := model.Photo{
			RoomId: int(room.ID),
			Url:    p,
		}
		if err = tx.Create(&photoModel).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	return tx.Commit().Error
}

func FindRoomById(id int) (room model.Room, err error) {
	err = global.DB.Where("id = ?", id).First(&room).Error
	return
}

func UpdateRoom(room model.Room, photo []string) (err error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Save(&room).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Unscoped().Where("room_id = ?", room.ID).Delete(&model.Photo{}).Error; err != nil {
		tx.Rollback()
		return
	}
	for _, p := range photo {
		photoModel := model.Photo{
			RoomId: int(room.ID),
			Url:    p,
		}
		if err = tx.Create(&photoModel).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	return tx.Commit().Error
}

func DeleteRoom(room model.Room) (err error) {
	err = global.DB.Unscoped().Delete(&room).Error
	return
}

func FindRoomPhoto(roomId int) (photos []model.Photo, err error) {
	err = global.DB.Where("room_id = ?", roomId).Find(&photos).Error
	return
}

func FindRoomList(page int64, pageSize int64, company int64) (rooms []model.Room, pages int64, totalCount int64, err error) {
	err = global.DB.Model(&model.Room{}).Where("company_id = ?", company).Count(&totalCount).Order("id desc").
		Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&rooms).Error
	pages = totalCount / pageSize
	if totalCount%(pageSize) != 0 {
		pages++
	}
	return
}
