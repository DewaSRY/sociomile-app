package utils

import (
	"encoding/json"
	"net/http"
)


type errorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Data interface{} `json:"data"`
}

type successResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := successResponse{
		Data: data,
		Code: statusCode,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := errorResponse{
		Error:   http.StatusText(statusCode),
		Code:    statusCode,
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}
