package handlers

import (
	"net/http"
	"strconv"

	"print-digital-proof-backend/models"
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

// CreateProofBody represents the request body for creating a new proof.
type CreateProofBody struct {
	OrderID       string `json:"orderId"`
	CustomerName  string `json:"customerName"`
	CustomerEmail string `json:"customerEmail"`
	ProductName   string `json:"productName"`
	FileName      string `json:"fileName"`
}

// CreateProof creates a new digital proof.
func (handler *ProofHandler) CreateProof(context *gin.Context) {
	var body CreateProofBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.OrderID == "" || body.CustomerName == "" || body.CustomerEmail == "" || body.ProductName == "" || body.FileName == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	proof := models.Proof{
		OrderID:       body.OrderID,
		CustomerName:  body.CustomerName,
		CustomerEmail: body.CustomerEmail,
		ProductName:   body.ProductName,
		FileName:      body.FileName,
		FileURL:       "/uploads/" + body.FileName,
		Comment:       "",
	}

	createdProof := handler.store.AddProof(proof)

	context.JSON(http.StatusCreated, createdProof)
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
