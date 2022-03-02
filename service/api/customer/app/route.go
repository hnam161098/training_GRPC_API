package app

import (
	"booking/protobuf/pb"
	"booking/service/api/customer/handler"
	"booking/service/transport"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Run() {
	customerConn, err := grpc.Dial(transport.CustomerGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	peopleClient := pb.NewCustomerServiceClient(customerConn)
	h := handler.NewCustomerHandler(peopleClient)
	g := gin.Default()

	// routes
	g.POST("/customer/create", h.CreateCustomer)
	g.PUT("/customer/update", h.UpdateCustomer)
	g.GET("/customer/find", h.FindCustomerById)
	fmt.Println("Customer API run port: 10000")
	g.Run(transport.CustomerAPI)
}
