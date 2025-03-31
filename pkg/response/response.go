package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // thông báo lỗi
	Data    interface{} `json:"data"`    // dữ liệu return
}

type ErrorResponseData struct {
	Code   int         `json:"code"`   // status code
	Err    string      `json:"error"`  // thong bao loi
	Detail interface{} `json:"detail"` // du lai return
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})

}
func ErrorResponse(c *gin.Context, code int, message string) {
	// message == "" set msg[code]
	if message == "" {
		message = msg[code]
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})

}
