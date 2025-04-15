package handlers

import (
	"go-gin/response"

	"github.com/gin-gonic/gin"
)

type healthcheck struct{}

func NewHealthcheck() *healthcheck {
	return &healthcheck{}
}

func (h *healthcheck) GetHealth(c *gin.Context) {
	res := "app is running well"

	response.Ok(c, res, "ok")
}
