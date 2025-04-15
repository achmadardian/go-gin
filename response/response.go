package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context, data interface{}, message ...string) {
	msg := ""

	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}

	c.JSON(http.StatusOK, res)
}

func NotFound(c *gin.Context, message ...string) {
	msg := ""

	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Code:    http.StatusNotFound,
		Message: msg,
	}

	c.JSON(http.StatusNotFound, res)
}

func BadRequest(c *gin.Context, message ...string) {
	msg := ""

	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	}

	c.JSON(http.StatusBadRequest, res)
}

func InternalServerError(c *gin.Context) {
	res := ApiResponse{
		Code: http.StatusInternalServerError,
	}

	c.JSON(http.StatusInternalServerError, res)
}
