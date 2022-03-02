package repositories

import (
	models "booking/service/grpc/booking/model"
	"context"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, model *models.BookingModel) (*models.BookingModel, error)
	FindBookingByCode(ctx context.Context, code string) (*models.BookingModel, error)
	CancelBooking(ctx context.Context, model *models.BookingModel) (*models.BookingModel, error)
}

func (m *dbManager) CreateBooking(ctx context.Context, model *models.BookingModel) (*models.BookingModel, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbManager) FindBookingByCode(ctx context.Context, code string) (*models.BookingModel, error) {
	var result models.BookingModel
	err := m.Where(&models.BookingModel{Code: code}).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *dbManager) CancelBooking(ctx context.Context, model *models.BookingModel) (*models.BookingModel, error) {
	err := m.Where(&models.BookingModel{Code: model.Code}).Updates(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
