package response

type AdditionalInfoErrorResponse struct {
	Key   string `json:"errorCode"`
	Value string `json:"message"`
}
