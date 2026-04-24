package calculator

// Zone represents a heart rate training zone.
// swagger:model
type Zone struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type Calculator struct {
	// Placeholder for calculator operations
}
