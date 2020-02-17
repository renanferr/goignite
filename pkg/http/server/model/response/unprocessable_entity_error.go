package response

type UnprocessableEntityErrorResponse struct {
	ErrorResponse
	ValidationErrors []ValidationErrorResponse `json:"validationErrors,omitempty"`
}
