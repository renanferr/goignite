package response

type ValidationError struct {
	FieldName       string `json:"fieldName"`
	RestrictionType string `json:"restrictionType"`
	Message         string `json:"message"`
}
