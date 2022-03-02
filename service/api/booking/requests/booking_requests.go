package requests

type BookingRequest struct {
	Code       string `json:"code"`
	Status     int64  `json:"status"`
	CustomerId string `json:"customer_id"`
}

type CancelBookingRequest struct {
	Code string `json:"code"`
}
