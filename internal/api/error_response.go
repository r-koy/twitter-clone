package api

type InnerError struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error InnerError `json:"error"`
}
