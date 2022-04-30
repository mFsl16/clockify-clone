package response

import "net/http"

type CommonRs struct {
	HttpStatus int         `json:"-"`
	Status     string      `json:"status"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func (rs CommonRs) SetSuccess(data interface{}) CommonRs {
	return CommonRs{
		HttpStatus: http.StatusOK,
		Status:     "success",
		Message:    "success",
		Data:       data,
	}
}

func (rs CommonRs) SetFailed(httpStatus int, message interface{}) CommonRs {
	status := "error"

	if httpStatus > 399 && httpStatus < 500 {
		status = "failed"
	}

	return CommonRs{
		HttpStatus: httpStatus,
		Status:     status,
		Message:    message,
	}
}
