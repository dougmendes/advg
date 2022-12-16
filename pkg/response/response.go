package response

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(data interface{}) Response {
	return Response{data, ""}
}

func DecodeError(err string) Response {
	return Response{nil, err}
}
