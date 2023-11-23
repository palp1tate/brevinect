package form

type BookRoomForm struct {
	RoomId    int64  `form:"roomId" json:"roomId" binding:"required"`
	StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
	EndTime   int64  `form:"endTime" json:"endTime" binding:"required"`
	Theme     string `form:"theme" json:"theme" binding:"required"`
}

type UpdateBookForm struct {
	BookId    int64  `form:"bookId" json:"bookId" binding:"required"`
	RoomId    int64  `form:"roomId" json:"roomId" binding:"required"`
	StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
	EndTime   int64  `form:"endTime" json:"endTime" binding:"required"`
	Theme     string `form:"theme" json:"theme" binding:"required"`
}
