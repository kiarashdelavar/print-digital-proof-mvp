package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Digital Proof API is running",
		})
	})

	router.Run(":8080")
}
