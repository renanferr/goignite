package girestresponse

type ValidationError struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}
