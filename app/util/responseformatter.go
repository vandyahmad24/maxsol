package util

type ResponseWithData struct {
	Success string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErrorWithData struct {
	Success string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"data"`
}

func ApiResponse(message string, data interface{}) interface{} {

	jsonResponse := ResponseWithData{
		Success: "Success",
		Message: message,
		Data:    data,
	}
	return jsonResponse

}

func ApiErrorResponse(message string, data interface{}) interface{} {

	jsonResponse := ResponseErrorWithData{
		Success: "Error",
		Message: message,
		Error:   data,
	}
	return jsonResponse

}
