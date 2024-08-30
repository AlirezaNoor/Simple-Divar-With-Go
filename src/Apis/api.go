package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServer(){
	r :=gin.New()
	r.Use(gin.logger(),gin.Recovery())
	v1 :=r.Group("/api/v1")
	{
		v1.GET("/healthChecker", func (c *gin.Context)  {
			c.JSON(http.StatusOK,"Working")
			return
		})
	}
r.Run("5005")
}