package main

import (
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/official/router"
)

func main() {
	gin.SetMode("debug")
	router := router.InitRouter()

	manners.ListenAndServe(":8080", router)
	// router.Run(":8080")
}
