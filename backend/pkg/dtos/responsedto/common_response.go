package responsedto


type ErrorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Error   string `json:"error,omitempty"`
}

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
}
