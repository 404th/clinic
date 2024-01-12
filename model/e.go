package model

type ErrorResponse struct {
	Message string
	Data    interface{}
}

type SuccessResponse struct {
	Message string
	Data    interface{}
}
