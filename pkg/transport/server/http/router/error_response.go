package router

type ErrorResponse struct {
	HttpStatusCode int                           `json:"httpStatusCode"`
	ErrorCode      string                        `json:"errorCode,omitempty"`
	Message        string                        `json:"message"`
	Info           string                        `json:"info,omitempty"`
	AdditionalInfo []AdditionalInfoErrorResponse `json:"additionalInfo,omitempty"`
}
