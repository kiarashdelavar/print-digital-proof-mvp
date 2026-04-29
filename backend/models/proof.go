package models

import "time"

// ProofStatus represents the current approval state of a digital proof.
type ProofStatus string

const (
	StatusWaitingForApproval ProofStatus = "WAITING_FOR_APPROVAL"
	StatusApproved           ProofStatus = "APPROVED"
	StatusChangesRequested   ProofStatus = "CHANGES_REQUESTED"
)

// Proof represents a digital proof that belongs to a customer order.
type Proof struct {
	ID            int         `json:"id"`
	OrderID       string      `json:"orderId"`
	CustomerName  string      `json:"customerName"`
	CustomerEmail string      `json:"customerEmail"`
	ProductName   string      `json:"productName"`
	FileName      string      `json:"fileName"`
	FileURL       string      `json:"fileUrl"`
	Status        ProofStatus `json:"status"`
	Comment       string      `json:"comment"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}
