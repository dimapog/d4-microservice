package calculator

type BMIResponse struct {
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	BMI      float64 `json:"bmi"`
	Category string  `json:"category"`
}

type HRZResponse struct {
	Age   int    `json:"age"`
	MaxHR int    `json:"max_hr"`
	Zones []Zone `json:"zones"`
}
