package response

type SuccessResponse struct {
	Lawyer interface{} `json:"lawyer,omitempty"`
}

type ErrorResponse struct {
	Error interface{} `json:"error, omitempty"`
}

func LaweryResponse(lawyer interface{}) SuccessResponse {
	return SuccessResponse{lawyer}
}

func DecodeError(err interface{}) ErrorResponse {
	return ErrorResponse{err}
}
