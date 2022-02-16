package handler

import (
	"booking/protobuf/pb"
	"booking/service/grpc/customer/repositories"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(customerRepo repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepo,
	}, nil
}
