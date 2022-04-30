package exception

import (
	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/model/response"
	"net/http"
)

func CustomHttpErrorHandler(err error, c echo.Context) {

	code := http.StatusInternalServerError
	message := err.Error()

	if commonException, ok := err.(CommonException); ok {
		c.Logger().Error(commonException)
		code = commonException.HttpStatus
		message = commonException.Message
	}

	commonRs := response.CommonRs{}
	c.JSON(code, commonRs.SetFailed(code, message))
}
