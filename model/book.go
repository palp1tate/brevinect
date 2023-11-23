package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	UserId    int    `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoomId    int    `gorm:"not null"`
	Room      Room   `gorm:"foreignKey:RoomId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StartTime int64  `gorm:"not null"`
	EndTime   int64  `gorm:"not null"`
	Theme     string `gorm:"not null"`
}
