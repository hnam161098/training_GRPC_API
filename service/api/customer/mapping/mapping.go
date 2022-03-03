package mapping

import (
	"booking/protobuf/pb"
	"booking/service/api/customer/requests"
	"booking/service/api/customer/responses"
)

func MappingCustomerRequestForCreateCustomer(in *requests.CustomerRequest) *pb.Customer {
	return &pb.Customer{
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}

func MappingCustomerResponseForCreateCustomer(in *pb.Customer) *responses.CustomerResponse {
	return &responses.CustomerResponse{
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}

func MappingCustomerRequestForUpdateCustomer(in *requests.CustomerUpdateRequest) *pb.Customer {
	return &pb.Customer{
		Id:       in.ID,
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}

func MappingCustomerResponseForUpdateCustomer(in *pb.Customer) *responses.CustomerResponse {
	return &responses.CustomerResponse{
		ID:       in.Id,
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}

func MappingCustomerRequestForFindCustomerById(in *requests.FindCustomerRequest) *pb.FindCustomerRequest {
	return &pb.FindCustomerRequest{
		Id: in.ID,
	}
}

func MappingCustomerResponseForFindCustomerById(in *pb.Customer) responses.CustomerResponse {
	return responses.CustomerResponse{
		ID:       in.Id,
		Name:     in.Name,
		Password: in.Password,
		Address:  in.Address,
		Email:    in.Email,
		Phone:    in.Phone,
	}
}
