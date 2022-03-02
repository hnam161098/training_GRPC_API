package handler

import (
	"booking/protobuf/pb"
	"booking/service/grpc/booking/repositories"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	BookingRepository repositories.BookingRepository
	CustomerClient    pb.CustomerServiceClient
}

func NewBookingHandler(bookingRepo repositories.BookingRepository, customerClient pb.CustomerServiceClient) (*BookingHandler, error) {
	return &BookingHandler{
		BookingRepository: bookingRepo,
		CustomerClient:    customerClient,
	}, nil
}
