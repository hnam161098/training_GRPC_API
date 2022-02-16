package handler

import (
	"booking/service/api/customer/mapping"
	"booking/service/api/customer/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	customerRequest := requests.CustomerRequest{}

	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})

			return
		}
	}
	req := mapping.MappingCustomerRequestForCreateCustomer(&customerRequest)

	result, err := h.customerClient.CreateCustomer(c.Request.Context(), req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": result,
	})

}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	customerRequest := requests.CustomerUpdateRequest{}
	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})

			return
		}
	}
	req := mapping.MappingCustomerRequestForUpdateCustomer(&customerRequest)

	res, err := h.customerClient.UpdateCustomer(c, req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	result := mapping.MappingCustomerResponseForUpdateCustomer(res)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": result,
	})
}

func (h *customerHandler) FindCustomerById(c *gin.Context) {
	customerRequest := requests.FindCustomerRequest{}

	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})

			return
		}
	}
	req := mapping.MappingCustomerRequestForFindCustomerById(&customerRequest)

	res, err := h.customerClient.FindCustomerById(c.Request.Context(), req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	result := mapping.MappingCustomerResponseForFindCustomerById(res)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": result,
	})
}
