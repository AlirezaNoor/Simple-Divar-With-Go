package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health1(c *gin.Context) {
	c.JSON(http.StatusOK, "healthcjhack1")
	return
}

func (h *HealthHandler) Health2(c *gin.Context) {
	c.JSON(http.StatusOK, "healthcheckForPost")
}

func (h *HealthHandler) Health3(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("healthcheckForId %s", id))
}
