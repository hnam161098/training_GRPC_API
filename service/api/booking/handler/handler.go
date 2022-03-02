package handler

import (
	"booking/protobuf/pb"

	"github.com/gin-gonic/gin"
)

type BookingHandler interface {
	CreateBookingHandler(c *gin.Context)
	FindBookingByCodeHandler(c *gin.Context)
	CancelBookingHandler(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BookingServiceClient
}

func NewBookingHandler(bookingClient pb.BookingServiceClient) bookingHandler {
	return bookingHandler{
		bookingClient: bookingClient,
	}
}
