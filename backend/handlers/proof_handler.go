package handlers

import (
	"net/http"
	"strconv"

	"print-digital-proof-backend/storage"

	"github.com/gin-gonic/gin"
)

// ProofHandler handles all proof-related API requests.
type ProofHandler struct {
	store *storage.MemoryStore
}

// NewProofHandler creates a new proof handler.
func NewProofHandler(store *storage.MemoryStore) *ProofHandler {
	return &ProofHandler{
		store: store,
	}
}

// GetAllProofs returns all proofs.
func (handler *ProofHandler) GetAllProofs(context *gin.Context) {
	proofs := handler.store.GetAllProofs()

	context.JSON(http.StatusOK, proofs)
}

// GetProofByID returns one proof by ID.
func (handler *ProofHandler) GetProofByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid proof ID",
		})
		return
	}

	proof, err := handler.store.GetProofByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Proof not found",
		})
		return
	}

	context.JSON(http.StatusOK, proof)
}

// ApproveProof approves a proof.
func (handler *ProofHandler) ApproveProof(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid proof ID",
		})
		return
	}

	proof, err := handler.store.ApproveProof(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Proof not found",
		})
		return
	}

	context.JSON(http.StatusOK, proof)
}

// RequestChangesBody represents the request body for requesting proof changes.
type RequestChangesBody struct {
	Comment string `json:"comment"`
}

// RequestChanges saves customer feedback and marks the proof as changes requested.
func (handler *ProofHandler) RequestChanges(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid proof ID",
		})
		return
	}

	var body RequestChangesBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.Comment == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Comment is required",
		})
		return
	}

	proof, err := handler.store.RequestChanges(id, body.Comment)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Proof not found",
		})
		return
	}

	context.JSON(http.StatusOK, proof)
}
