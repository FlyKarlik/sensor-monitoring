package response

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, msg string, status int) {
	c.JSON(status, response{
		Code:    status,
		Message: &msg,
	})
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, response{
		Code: status,
		Data: data,
	})
}
