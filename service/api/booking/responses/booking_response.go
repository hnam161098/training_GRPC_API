package responses

import (
	"booking/service/api/customer/responses"
)

type BookingResponse struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Status     int64  `json:"status"`
	CustomerId string `json:"customer_id"`
}

type FindBookingByCodeResponse struct {
	BookingResponse
	CustomerInfor responses.CustomerResponse
}
