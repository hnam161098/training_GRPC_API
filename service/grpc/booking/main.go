package main

import (
	"booking/protobuf/pb"
	"booking/service/grpc/booking/handler"
	"booking/service/grpc/booking/repositories"
	"booking/service/transport"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	customerConnect, err := grpc.Dial(transport.CustomerGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	customerClient := pb.NewCustomerServiceClient(customerConnect)

	listen, err := net.Listen("tcp", transport.BookingPortGRPC)
	if err != nil {
		log.Fatal(err)
	}
	bookingRepos, err := repositories.NewDBManager()
	if err != nil {
		log.Fatal(err)
	}

	h, err := handler.NewBookingHandler(bookingRepos, customerClient)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterBookingServiceServer(s, h)
	fmt.Println("Booking GRPC run port: 9010")
	s.Serve(listen)
}
