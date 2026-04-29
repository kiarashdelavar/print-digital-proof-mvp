package storage

import (
	"errors"
	"sync"
	"time"

	"print-digital-proof-backend/models"
)

// MemoryStore keeps proof data in memory for the MVP.
// This is simple and fast for the first prototype.
// Later, this can be replaced by a real database.
type MemoryStore struct {
	proofs []models.Proof
	nextID int
	mu     sync.Mutex
}

// NewMemoryStore creates a new memory store with example proof data.
func NewMemoryStore() *MemoryStore {
	now := time.Now()

	return &MemoryStore{
		nextID: 3,
		proofs: []models.Proof{
			{
				ID:            1,
				OrderID:       "ORD-1001",
				CustomerName:  "Kiarash Delavar",
				CustomerEmail: "customer@example.com",
				ProductName:   "Business Cards",
				FileName:      "business-card-proof.pdf",
				FileURL:       "/uploads/business-card-proof.pdf",
				Status:        models.StatusWaitingForApproval,
				Comment:       "",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			{
				ID:            2,
				OrderID:       "ORD-1002",
				CustomerName:  "Demo Customer",
				CustomerEmail: "demo@example.com",
				ProductName:   "Flyer A5",
				FileName:      "flyer-proof.pdf",
				FileURL:       "/uploads/flyer-proof.pdf",
				Status:        models.StatusApproved,
				Comment:       "",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
		},
	}
}

// GetAllProofs returns all digital proofs.
func (store *MemoryStore) GetAllProofs() []models.Proof {
	store.mu.Lock()
	defer store.mu.Unlock()

	return store.proofs
}

// GetProofByID returns one proof by ID.
func (store *MemoryStore) GetProofByID(id int) (*models.Proof, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	for index := range store.proofs {
		if store.proofs[index].ID == id {
			return &store.proofs[index], nil
		}
	}

	return nil, errors.New("proof not found")
}

// AddProof adds a new proof to the store.
func (store *MemoryStore) AddProof(proof models.Proof) models.Proof {
	store.mu.Lock()
	defer store.mu.Unlock()

	now := time.Now()

	proof.ID = store.nextID
	proof.Status = models.StatusWaitingForApproval
	proof.CreatedAt = now
	proof.UpdatedAt = now

	store.proofs = append(store.proofs, proof)
	store.nextID++

	return proof
}

// ApproveProof changes the proof status to approved.
func (store *MemoryStore) ApproveProof(id int) (*models.Proof, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	for index := range store.proofs {
		if store.proofs[index].ID == id {
			store.proofs[index].Status = models.StatusApproved
			store.proofs[index].Comment = ""
			store.proofs[index].UpdatedAt = time.Now()
			return &store.proofs[index], nil
		}
	}

	return nil, errors.New("proof not found")
}

// RequestChanges changes the proof status and saves the customer's feedback.
func (store *MemoryStore) RequestChanges(id int, comment string) (*models.Proof, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	for index := range store.proofs {
		if store.proofs[index].ID == id {
			store.proofs[index].Status = models.StatusChangesRequested
			store.proofs[index].Comment = comment
			store.proofs[index].UpdatedAt = time.Now()
			return &store.proofs[index], nil
		}
	}

	return nil, errors.New("proof not found")
}
