package repositories

import (
	"booking/service/conf"
	models "booking/service/grpc/booking/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbManager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := gorm.Open(postgres.Open(conf.DB_URL))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.BookingModel{})

	return &dbManager{db}, nil
}
