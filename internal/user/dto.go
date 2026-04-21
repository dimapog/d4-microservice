package user

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Age              *int     `json:"age"`
	Gender           *string  `json:"gender"`
	Weight           *float64 `json:"weight"`
	Height           *float64 `json:"height"`
	RestingHeartRate *int     `json:"resting_heart_rate"`
	Units            *string  `json:"units"`
}

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

type GetUserRequest struct {
	ID string `uri:"id" binding:"required"`
}
