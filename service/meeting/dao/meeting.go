package dao

import (
	"github.com/palp1tate/brevinect/model"
	"github.com/palp1tate/brevinect/service/meeting/global"
	"github.com/palp1tate/brevinect/util"
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

func CreateBook(book *model.Book) (err error) {
	err = global.DB.Create(&book).Error
	return
}

func FindBookList(page int64, pageSize int64, userId int64) (books []model.Book, pages int64, totalCount int64, err error) {
	err = global.DB.Model(&model.Book{}).Where("user_id = ?", userId).Count(&totalCount).Order("id desc").
		Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&books).Error
	pages = totalCount / pageSize
	if totalCount%(pageSize) != 0 {
		pages++
	}
	return
}

func FindBookById(bookId int, userId int) (book model.Book, err error) {
	err = global.DB.Where("id = ? and user_id = ?", bookId, userId).First(&book).Error
	return
}

func DeleteBook(book model.Book) (err error) {
	err = global.DB.Unscoped().Delete(&book).Error
	return
}

func UpdateBook(book model.Book) (err error) {
	err = global.DB.Save(&book).Error
	return
}

func FindBook(roomId int) (book []model.Book, err error) {
	todayZeroTime := util.GetTodayZeroTime()
	t := todayZeroTime + 2*24*60*60
	err = global.DB.Where("room_id = ? and start_time >= ? and end_time < ?",
		roomId, todayZeroTime, t).Find(&book).Error
	return
}

func FindBooks(userId int) (books []model.Book, err error) {
	err = global.DB.Where("user_id = ?", userId).Find(&books).Error
	return
}

func FindUserById(userId int) (user model.User, err error) {
	err = global.DB.Where("id = ?", userId).First(&user).Error
	return
}
