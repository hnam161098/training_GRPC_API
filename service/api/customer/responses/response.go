package responses

type CustomerResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    int64  `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}
