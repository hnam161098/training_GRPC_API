package handler

import (
	"booking/protobuf/pb"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	CreatCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	FindCustomerById(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.CustomerServiceClient
}

func NewCustomerHandler(customerClient pb.CustomerServiceClient) customerHandler {
	return customerHandler{
		customerClient: customerClient,
	}
}
