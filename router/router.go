package router

import (
	_ "github.com/official/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	//"github.com/official/handler"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	// r.GET("/login", api.GetAuth)

	r.Use(gin.Recovery())

	// r.GET("/services/:id", controller.OneGoods)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.Use()
	{

		// Router.POST("/graphql", graphql.GraphqlHandler())
		// Router.GET("/graphql", graphql.GraphqlHandler())
		// r.GET("/index", controller.Index)
		// r.GET("/services", controller.MultiGoods)
		// api.PUT("/services/:id", controller.Update)
		// api.DELETE("/services/:id", controller.DeleteGoods)

	}

	return r
}
