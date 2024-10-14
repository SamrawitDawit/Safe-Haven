package utils

import (
	"log"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(status int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(status int, message string, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Error:   err,
	}
}

func LogError(message string, err error) {
	if err != nil {
		log.Printf("[ERROR] %s: %v\n", message, err)
	} else {
		log.Printf("[ERROR] %s\n", message)
	}
}
