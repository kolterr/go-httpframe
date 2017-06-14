package protocol

type (
	ErrorResponse struct {
		Code  int    `json:"code"`
		Error string `json:"details"`
	}
)
