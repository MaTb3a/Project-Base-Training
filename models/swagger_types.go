package models

// SwaggerErrorResponse represents an error response using gin.H
// @name SwaggerErrorResponse
type SwaggerErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SwaggerSuccessResponse represents a success response using gin.H
// @name SwaggerSuccessResponse
type SwaggerSuccessResponse struct {
	Message string `json:"message" example:"success message"`
}