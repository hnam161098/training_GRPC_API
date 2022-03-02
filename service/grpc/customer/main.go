package main

import (
	"booking/protobuf/pb"
	"booking/service/grpc/customer/handler"
	"booking/service/grpc/customer/repositories"
	"booking/service/transport"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", transport.CustomerGRPC)
	if err != nil {
		log.Fatal(err)
	}
	customerRepos, err := repositories.NewDBManager()
	if err != nil {
		log.Fatal(err)
	}
	h, err := handler.NewCustomerHandler(customerRepos)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)
	fmt.Println("Customer GRPC run port: 9000")
	s.Serve(listen)
}
