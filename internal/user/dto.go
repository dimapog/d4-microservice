package user

// CreateUserRequest represents payload for creating a new user.
// swagger:model
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserRequest represents payload for updating an authenticated user.
// swagger:model
type UpdateUserRequest struct {
	Age              *int     `json:"age"`
	Gender           *string  `json:"gender"`
	Weight           *float64 `json:"weight"`
	Height           *float64 `json:"height"`
	RestingHeartRate *int     `json:"resting_heart_rate"`
	Units            *string  `json:"units"`
}

// UserResponse represents the returned user profile.
// swagger:model
type UserResponse struct {
	ID               uint     `json:"id"`
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Age              *int     `json:"age,omitempty"`
	Gender           *string  `json:"gender,omitempty"`
	Weight           *float64 `json:"weight,omitempty"`
	Height           *float64 `json:"height,omitempty"`
	RestingHeartRate *int     `json:"resting_heart_rate,omitempty"`
	Units            *string  `json:"units,omitempty"`
}

// ErrorResponse represents a standard JSON error payload.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

type GetUserRequest struct {
	ID string `uri:"id" binding:"required"`
}
