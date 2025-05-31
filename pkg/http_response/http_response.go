package http_response

type Response struct {
	Success bool `json:"success"`
	Message any  `json:"message,omitempty"`
	Data    any  `json:"data"`
	Meta    any  `json:"meta,omitempty"`
}
