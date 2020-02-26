package op5

type APIResponseError struct {
	Error     string `json:"error"`
	FullError string `json:"full_error"`
}
