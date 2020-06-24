package controller

import "github.com/gin-gonic/gin"

func SetRoutes(route *gin.Engine) *gin.Engine {
	v := route.Group("/api")
	v.Use(gin.Logger())
	{
		v.POST("/check_live", CheckLive)
	}
	return route
}
