package api

import (
	"SimpleProjectWithGo/api/routers"
	"SimpleProjectWithGo/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	api:=r.Group("/api")
	v1 := api.Group("/api")
	{
		health := v1.Group("/health1")
		routers.HealthRouter(health)

	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
