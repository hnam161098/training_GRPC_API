package repositories

import (
	models "booking/service/grpc/booking/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbManager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=12345 dbname=booking port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.BookingModel{})

	return &dbManager{db}, nil
}
