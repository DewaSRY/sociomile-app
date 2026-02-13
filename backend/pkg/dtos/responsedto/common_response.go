package responsedto

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error,omitempty"`
}

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type PaginateMetaData struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
