package responder

// swagger:model
type ErrorResponse struct {
	// error message
	Message string `json:"message"`
}
