package handler

import (
	"booking/protobuf/pb"
	"booking/service/grpc/booking/mapping"
	"booking/service/grpc/common"
	"context"
)

func (h *BookingHandler) CreateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	req := mapping.MappingBookingModelForCreateBooking(in)

	booking, err := h.BookingRepository.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	result := mapping.MappingBookingModel(booking)
	return result, nil
}

func (h *BookingHandler) FindBookingByCode(ctx context.Context, in *pb.BookingRequest) (*pb.BookingInformation, error) {
	bookingInfor, err := h.BookingRepository.FindBookingByCode(ctx, in.Code)
	if err != nil {
		return nil, err
	}

	customerInfor, err := h.CustomerClient.FindCustomerById(ctx, &pb.FindCustomerRequest{
		Id: bookingInfor.Customer_id,
	})
	if err != nil {
		return nil, err
	}
	return mapping.MappingBookingInformation(bookingInfor, customerInfor), nil

}

func (h *BookingHandler) CancelBooking(ctx context.Context, in *pb.CancelRequest) (*pb.Booking, error) {
	bookingCancel, err := h.BookingRepository.FindBookingByCode(ctx, in.Code)
	if err != nil {
		return nil, err
	}
	bookingCancel.Status = common.InActiveStatus

	newBooking, err := h.BookingRepository.CancelBooking(ctx, bookingCancel)
	if err != nil {
		return nil, err
	}

	// return &pb.Booking{
	// 	Id:         newBooking.ID,
	// 	Status:     newBooking.Status,
	// 	CustomerId: newBooking.Customer_id,
	// 	Code:       newBooking.Code,
	// }, nil
	result := mapping.MappingBookingModel(newBooking)
	return result, nil
}
