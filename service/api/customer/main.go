package main

import (
	"booking/protobuf/pb"
	"booking/service/api/customer/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	customerConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	peopleClient := pb.NewCustomerServiceClient(customerConn)
	h := handler.NewCustomerHandler(peopleClient)
	g := gin.Default()

	// routes
	g.POST("/create", h.CreateCustomer)
	g.PUT("/update", h.UpdateCustomer)
	g.GET("/find", h.FindCustomerById)
	fmt.Println("csutomer run port: 8080")
	g.Run(":8080")
}
