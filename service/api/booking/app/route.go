package app

import (
	"booking/protobuf/pb"
	"booking/service/api/booking/handler"
	"booking/service/transport"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Run() {
	bookingConnect, err := grpc.Dial(transport.BookingPortGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	bookingClient := pb.NewBookingServiceClient(bookingConnect)
	h := handler.NewBookingHandler(bookingClient)
	g := gin.Default()

	//creat routes
	g.POST("/booking/create", h.CreateBookingHandler)
	g.POST("/booking/cancel", h.CancelBookingHandler)
	g.POST("/booking/view", h.FindBookingByCodeHandler)
	fmt.Println("Booking API run port: 9010")
	g.Run(transport.BookingAPI)
}
