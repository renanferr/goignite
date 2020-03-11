package rest

type ValidationErrorResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}
