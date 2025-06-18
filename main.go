package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santori/ping-checker/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/check", handlers.CheckReq)

	r.Run(":8080")
}
