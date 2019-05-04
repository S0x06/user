package service

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	// 返回JSON数据
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
