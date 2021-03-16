package girestresponse

type AdditionalInfoError struct {
	Key   string `json:"errorCode"`
	Value string `json:"message"`
}
