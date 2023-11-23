package model

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	RoomId int    `gorm:"not null"`
	Room   Room   `gorm:"foreignKey:RoomId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Url    string `gorm:"not null"`
}
