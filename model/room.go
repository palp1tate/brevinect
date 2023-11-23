package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	CompanyId int     `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyId;AssociationForeignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Name     string `gorm:"not null"`
	Capacity int    `gorm:"not null"`
	Location string `gorm:"not null"`
	Facility string `gorm:"not null"`
}
