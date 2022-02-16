package handler

import (
	"booking/protobuf/pb"
	"booking/service/grpc/customer/mapping"
	"context"
)

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	req := mapping.MappingCustomerModelForCreateCustomer(in)

	customer, err := h.customerRepository.CreateCustomer(ctx, req)
	if err != nil {
		return nil, err
	}
	result := mapping.MappingCustomerModel(customer)
	return result, nil
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	req := mapping.MappingCustomerModelForUpdateCustomer(in)

	customer, err := h.customerRepository.UpdateCustomer(ctx, req)
	if err != nil {
		return nil, err
	}
	result := mapping.MappingCustomerModel(customer)
	return result, nil
}

func (h *CustomerHandler) FindCustomerById(ctx context.Context, in *pb.FindCustomerRequest) (*pb.Customer, error) {
	req := mapping.MappingCustomerModelForFindCustomer(in)

	customer, err := h.customerRepository.FindCustomerById(ctx, req)
	if err != nil {
		return nil, err
	}
	result := mapping.MappingCustomerModel(customer)
	return result, nil
}
