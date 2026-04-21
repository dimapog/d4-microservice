package ai

type PersonalCalculationRequest struct {
	Age              int     `json:"age" binding:"required"`
	Gender           string  `json:"gender" binding:"required"`
	Weight           float64 `json:"weight" binding:"required"`
	Height           float64 `json:"height" binding:"required"`
	RestingHeartRate int     `json:"resting_heart_rate" binding:"required"`
	Units            string  `json:"units" binding:"required"`
}

type PersonalCalculationResponse struct {
	Response interface{} `json:"response"`
}
