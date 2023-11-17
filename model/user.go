package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"not null"`
	Password  string  `gorm:"not null"`
	Mobile    string  `gorm:"not null"`
	CompanyId int     `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Avatar    string  `gorm:"not null;default:http://s42es6gy4.hn-bkt.clouddn.com/avatar.jpg"`
	Role      int     `gorm:"not null;default:1;check: role in(1,2)"` // 1:普通用户 2:管理员
}
