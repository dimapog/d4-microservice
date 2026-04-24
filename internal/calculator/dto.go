package calculator

// BMIResponse represents the computed BMI and category.
// swagger:model
type BMIResponse struct {
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	BMI      float64 `json:"bmi"`
	Category string  `json:"category"`
}

// HRZResponse represents calculated heart rate zones.
// swagger:model
type HRZResponse struct {
	Age   int    `json:"age"`
	MaxHR int    `json:"max_hr"`
	Zones []Zone `json:"zones"`
}

// ErrorResponse represents a standard JSON error payload.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}
