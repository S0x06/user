package router

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	// r.GET("/login", api.GetAuth)

	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	api.Use(jwt.JWT())
	{

	}

	return r
}
