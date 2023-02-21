package models

const ResponseStatusFailed = "failed"
const ResponseStatusSuccess = "success"
const ResponseStatusFailedInternal = "internal server error"

var ResponseStatusFailedMap = map[string]interface{}{
	"status":  ResponseStatusFailed,
	"message": ResponseStatusFailedInternal,
}

type Response struct {
	Code int
	Body map[string]interface{}
}

func NewResponse(code int, body map[string]interface{}) *Response {
	return &Response{
		Code: code,
		Body: body,
	}
}
