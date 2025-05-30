package http_response

type Response struct {
	Success bool   `json:"success"`
	Info    string `json:"info,omitempty"`
	Error   error  `json:"error,omitempty"`
	Data    any    `json:"data"`
	Meta    any    `json:"meta,omitempty"`
}
