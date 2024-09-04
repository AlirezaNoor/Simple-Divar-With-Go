package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type health struct {}




func NewHealthInstance () *health {

	return &health{}
}

func (h *health) HealthCheck(c *gin.Context){

c.JSON(http.StatusOK,"health")
}