package request

// CreateUser represents request object of person using this service.
type CreateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
