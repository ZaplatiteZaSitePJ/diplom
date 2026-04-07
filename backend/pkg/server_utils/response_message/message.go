package response_message

type Response struct {
	StatusCode int  `json:"status_code"`
	IsError    bool `json:"is_error"`
	Data       any  `json:"data"`
}

func NewResponseMessage(sc int, d any, is_err bool) *Response {
	return &Response{
		StatusCode: sc,
		Data:       d,
		IsError:    is_err,
	}
}
