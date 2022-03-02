package handler

import (
	"booking/protobuf/pb"
	"booking/service/api/booking/mapping"
	"booking/service/api/booking/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *bookingHandler) CreateBookingHandler(c *gin.Context) {
	bookingRequest := requests.BookingRequest{}

	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
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

	req := mapping.MappingBookingRequestForCreateBooking(&bookingRequest)

	result, err := h.bookingClient.CreateBooking(c.Request.Context(), req)

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

func (h *bookingHandler) FindBookingByCodeHandler(c *gin.Context) {
	bookingRequest := requests.BookingRequest{}
	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
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

	req := &pb.BookingRequest{
		Code: bookingRequest.Code,
	}

	res, err := h.bookingClient.FindBookingByCode(c.Request.Context(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	result := mapping.MappingBookingResponseForFindBookingByCode(res)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": result,
	})
}

func (h *bookingHandler) CancelBookingHandler(c *gin.Context) {
	cancelBookingRequest := requests.CancelBookingRequest{}

	if err := c.ShouldBindJSON(&cancelBookingRequest); err != nil {
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

	req := mapping.MappingBookingRequestForCancelBooking(&cancelBookingRequest)

	result, err := h.bookingClient.CancelBooking(c.Request.Context(), req)

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
