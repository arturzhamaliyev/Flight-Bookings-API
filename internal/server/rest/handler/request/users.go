package request

// SignUp represents request object of person who wants to sign up.
type SignUp struct {
	Phone    *string `json:"phone,omitempty"`
	Password string  `json:"password" binding:"required"`
	Email    string  `json:"email" binding:"required"`
}

// SignIn represents request object of person who wants to sign in.
type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
