package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type health struct{}

func NewHealthInstance() *health {

	return &health{}
}

func (h *health) HealthCheck(c *gin.Context) {

	c.JSON(http.StatusOK, "health")
}
func (h *health) HealthCheckPost(c *gin.Context) {

	c.JSON(http.StatusOK, "health")
}

func (h *health) HealthCheckById(c *gin.Context) {

	c.JSON(http.StatusOK, fmt.Sprintf("health %s", c.Params.ByName("id")))
}
