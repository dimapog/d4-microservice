package csv

// CSVUploadResponse contains the async import job metadata.
// swagger:model
type CSVUploadResponse struct {
	JobID  string `json:"job_id"`
	Status string `json:"status"`
}

// ErrorResponse represents a standard JSON error payload.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}
