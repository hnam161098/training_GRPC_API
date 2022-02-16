package requests

type CustomerRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    int64  `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type CustomerUpdateRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    int64  `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type FindCustomerRequest struct {
	ID string `json:"id" binding:"required"`
}
