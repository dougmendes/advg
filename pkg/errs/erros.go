package err

func ErrorResponse(code int, message string, detail error) *Error {
	errorResponse := Error{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
	return &errorResponse
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  error  `json:"detail"`
}
