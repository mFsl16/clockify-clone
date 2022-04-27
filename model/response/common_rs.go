package response

import "net/http"

type CommonRs struct {
	Status  int         `json: "status"`
	Message interface{} `json: "message"`
	Data    interface{} `json: "data"`
}

func (rs CommonRs) SetSuccess(data interface{}) CommonRs {
	return CommonRs{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	}
}
