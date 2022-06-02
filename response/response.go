package response

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func BuildResponse(status bool, message string, data interface{}) *Response {
	return &Response{
		Status: status,
		Message: message,
		Data: data,
	}
}

func BuildErrorResponse(message string) *Response {
	return &Response{
		Status: false,
		Message: message,
		Data: nil,
	}
}