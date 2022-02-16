package mapping

import (
	"booking/protobuf/pb"
	models "booking/service/grpc/customer/model"

	"github.com/google/uuid"
)

func MappingCustomerModel(in *models.CustomerModel) *pb.Customer {
	return &pb.Customer{
		Id:       in.ID,
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Phone:    in.Phone,
		Email:    in.Email,
	}
}

func MappingCustomerModelForCreateCustomer(in *pb.Customer) *models.CustomerModel {
	return &models.CustomerModel{
		ID:       uuid.New().String(),
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Phone:    in.Phone,
		Email:    in.Email,
	}
}

func MappingCustomerModelForUpdateCustomer(in *pb.Customer) *models.CustomerModel {
	return &models.CustomerModel{
		ID:       in.Id,
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Phone:    in.Phone,
		Email:    in.Email,
	}
}

func MappingCustomerModelForFindCustomer(in *pb.FindCustomerRequest) *models.CustomerModel {
	return &models.CustomerModel{
		ID: in.Id,
	}
}
