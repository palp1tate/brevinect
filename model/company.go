package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Address        string `gorm:"not null"`
	OfficialMobile string `gorm:"not null"`
	OfficialSite   string `gorm:"not null"`
	CompanyType    string `gorm:"not null"`
	Introduction   string `gorm:"not null"`
	Picture        string `gorm:"not null"`
}
