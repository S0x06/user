package main

import (
	"github.com/user/router"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode("debug")
	router := router.InitRouter()
	router.Run(":80")
}
