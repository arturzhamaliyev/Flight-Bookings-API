package request

// CreateUser represents request object of person using this service.
type CreateUser struct {
	Phone    *string `json:"phone,omitempty"`
	Password string  `json:"password" binding:"required"`
	Email    string  `json:"email" binding:"required"`
}
