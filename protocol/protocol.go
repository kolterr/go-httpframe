package protocol

type ErrorWrapper struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
}
