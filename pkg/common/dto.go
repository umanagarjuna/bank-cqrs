package common

// BaseResponse is a standard response structure.
type BaseResponse struct {
	Message string `json:"message"`
}

// OpenAccountResponse is the response for opening a new account.
type OpenAccountResponse struct {
	BaseResponse
	ID string `json:"id"`
}
