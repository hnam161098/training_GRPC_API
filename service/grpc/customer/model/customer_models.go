package models

import "gorm.io/gorm"

type CustomerModel struct {
	gorm.Model
	ID       string `gorm:"type:varchar(256)"`
	Name     string `gorm:"type:varchar(256)"`
	Password string `gorm:"type:varchar(50)"`
	Address  string `gorm:"type:varchar(256)"`
	Phone    int64  `gorm:"type:int"`
	Email    string `gorm:"type:varchar(50)"`
}
