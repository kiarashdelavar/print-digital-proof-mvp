package main

import (
	"net/http"

	"print-digital-proof-backend/handlers"
	"print-digital-proof-backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
