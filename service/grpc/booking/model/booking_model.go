package models

import "gorm.io/gorm"

type BookingModel struct {
	gorm.Model
	ID          string `gorm:"type:varchar(256)"`
	Code        string `gorm:"type:varchar(256)"`
	Status      int64  `gorm:"type:int"`
	Customer_id string `gorm:"type:varchar(256)"`
}
