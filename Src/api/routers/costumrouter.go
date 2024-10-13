package routers

import (
	"SimpleProjectWithGo/api/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {

	handler := handlers.NewHealthHandler()
	check:=handlers.HealthHandler

	r.GET("/", handler.Health1)
	r.POST("/check2", handler.Health2)
	r.GET("/byid:id", handler.Health3)

}
