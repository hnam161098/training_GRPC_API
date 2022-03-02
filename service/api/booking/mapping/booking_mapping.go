package mapping

import (
	"booking/protobuf/pb"
	"booking/service/api/booking/requests"
	"booking/service/api/booking/responses"
	cus "booking/service/api/customer/responses"
)

func MappingBookingRequestForCreateBooking(in *requests.BookingRequest) *pb.Booking {
	return &pb.Booking{
		CustomerId: in.CustomerId,
	}
}

func MappingBookingResponseForCreateBooking(in *pb.Booking) *responses.BookingResponse {
	return &responses.BookingResponse{
		ID:         in.Id,
		Code:       in.Code,
		Status:     in.Status,
		CustomerId: in.CustomerId,
	}
}

func MappingBookingResponseForFindBookingByCode(res *pb.BookingInformation) *responses.FindBookingByCodeResponse {
	return &responses.FindBookingByCodeResponse{
		BookingResponse: responses.BookingResponse{
			ID:         res.BookingDetail.Id,
			Code:       res.BookingDetail.Code,
			CustomerId: res.BookingDetail.CustomerId,
			Status:     res.BookingDetail.Status,
		},
		CustomerInfor: cus.CustomerResponse{
			ID:       res.CustomerDetail.Id,
			Name:     res.CustomerDetail.Name,
			Password: res.CustomerDetail.Password,
			Phone:    res.CustomerDetail.Phone,
			Address:  res.CustomerDetail.Address,
			Email:    res.CustomerDetail.Email,
		},
	}
}

func MappingBookingRequestForCancelBooking(in *requests.CancelBookingRequest) *pb.CancelRequest {
	return &pb.CancelRequest{
		Code: in.Code,
	}
}
