package routers

import (
	handlers "project/Apis/Handlers"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	h := handlers.NewHealthInstance()

	r.GET("/", h.HealthCheck)
	r.POST("/", h.HealthCheckPost)
	r.GET("/:id", h.HealthCheckById)
}
