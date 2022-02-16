package repositories

import (
	models "booking/service/grpc/customer/model"
	"context"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, model *models.CustomerModel) (*models.CustomerModel, error)
	UpdateCustomer(ctx context.Context, model *models.CustomerModel) (*models.CustomerModel, error)
	FindCustomerById(ctx context.Context, model *models.CustomerModel) (*models.CustomerModel, error)
}

func (m *dbManager) CreateCustomer(ctx context.Context, model *models.CustomerModel) (*models.CustomerModel, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbManager) UpdateCustomer(ctx context.Context, model *models.CustomerModel) (*models.CustomerModel, error) {
	err := m.Where(&models.CustomerModel{ID: model.ID}).Updates(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbManager) FindCustomerById(ctx context.Context, model *models.CustomerModel) (*models.CustomerModel, error) {
	var result models.CustomerModel
	err := m.Where(&models.CustomerModel{ID: model.ID}).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
