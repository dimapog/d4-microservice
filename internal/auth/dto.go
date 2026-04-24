package auth

// LoginRequest represents user login payload.
// swagger:model
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse contains the JWT returned after successful authentication.
// swagger:model
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse represents a standard JSON error payload.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}
