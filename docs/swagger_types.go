package docs

// ErrorResponse represents an error response using gin.H
// @name ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse represents a success response using gin.H
// @name SuccessResponse
type SuccessResponse struct {
	Message string `json:"message" example:"success message"`
}