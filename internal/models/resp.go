package models

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
}
