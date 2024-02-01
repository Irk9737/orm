package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 应答 Client 请求
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (res *Response) HandleResponse(c *gin.Context, code int, message string, data interface{}) {
	res.Code = code
	res.Message = message
	res.Data = data
	c.JSON(http.StatusOK, res)
}
