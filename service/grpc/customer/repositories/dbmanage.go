package repositories

import (
	"booking/service/conf"
	models "booking/service/grpc/customer/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbManager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := gorm.Open(postgres.Open(conf.DB_URL))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.CustomerModel{})

	return &dbManager{db}, nil
}
