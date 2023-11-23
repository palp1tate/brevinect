package dao

import (
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/service/meeting/global"
)

func FindRoomById(id int) (room model.Room, err error) {
	err = global.DB.Where("id = ?", id).First(&room).Error
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
