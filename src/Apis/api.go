package api

import (
 	"project/Apis/routers"

	"github.com/gin-gonic/gin"
)

func InternalServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		// v1.GET("/healthChecker", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, "Working")
		// 	return
		// })
			health:=v1.Group("/health")

			routers.Health(health) 
	}
	r.Run(":5005")
}
