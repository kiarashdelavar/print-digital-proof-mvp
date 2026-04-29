package main

import (
	"net/http"
	"time"

	"print-digital-proof-backend/handlers"
	"print-digital-proof-backend/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	memoryStore := storage.NewMemoryStore()
	proofHandler := handlers.NewProofHandler(memoryStore)

	router.GET("/api/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Digital Proof API is running",
		})
	})

	api := router.Group("/api")
	{
		api.GET("/proofs", proofHandler.GetAllProofs)
		api.GET("/proofs/:id", proofHandler.GetProofByID)
		api.POST("/proofs/:id/approve", proofHandler.ApproveProof)
		api.POST("/proofs/:id/request-changes", proofHandler.RequestChanges)
	}

	router.Run(":8080")
}
