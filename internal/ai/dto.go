package ai

// PersonalCalculationRequest contains the user health data for AI analysis.
// swagger:model
type PersonalCalculationRequest struct {
	Age              int     `json:"age" binding:"required"`
	Gender           string  `json:"gender" binding:"required"`
	Weight           float64 `json:"weight" binding:"required"`
	Height           float64 `json:"height" binding:"required"`
	RestingHeartRate int     `json:"resting_heart_rate" binding:"required"`
	Units            string  `json:"units" binding:"required"`
}

// PersonalCalculationResponse wraps the AI response payload.
// swagger:model
type PersonalCalculationResponse struct {
	Response interface{} `json:"response"`
}

// ErrorResponse represents a standard JSON error payload.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}
