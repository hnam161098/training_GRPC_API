package mapping

import (
	"booking/protobuf/pb"
	models "booking/service/grpc/booking/model"
	"booking/service/grpc/common"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

func MappingBookingModel(in *models.BookingModel) *pb.Booking {
	return &pb.Booking{
		Id:         in.ID,
		CustomerId: in.Customer_id,
		Code:       in.Code,
		Status:     in.Status,
	}
}

func MappingBookingModelForCreateBooking(in *pb.Booking) *models.BookingModel {
	return &models.BookingModel{
		ID:          uuid.New().String(),
		Code:        xid.New().String(),
		Customer_id: in.CustomerId,
		Status:      common.ActiveStatus,
	}
}

func MappingBookingModelForFindBookingByCode(in *pb.Booking) *models.BookingModel {
	return &models.BookingModel{
		Code: in.Code,
	}
}

func MappingBookingInformation(bookingModel *models.BookingModel, customerModel *pb.Customer) *pb.BookingInformation {
	return &pb.BookingInformation{
		CustomerDetail: customerModel,
		BookingDetail: &pb.Booking{
			Id:         bookingModel.ID,
			CustomerId: bookingModel.Customer_id,
			Code:       bookingModel.Code,
			Status:     bookingModel.Status,
		},
	}
}
