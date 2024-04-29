package responses

import (
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func CreatedResponse(data interface{}) *Response {
	return &Response{
		Message: "Successfully Created",
		Status:  http.StatusCreated,
		Data:    data,
	}
}

func ValidationErrorResponse(data interface{}) *Response {
	return &Response{
		Message: "Validation Error",
		Status:  http.StatusBadRequest,
		Data:    data,
	}
}
