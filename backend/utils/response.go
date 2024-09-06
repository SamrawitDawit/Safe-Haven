package utils

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err string) Response {
	return Response{
		Message: message,
		Error:   err,
	}
}
