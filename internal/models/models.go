package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a client or freelancer
type User struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	Email        string     `db:"email" json:"email"`
	Name         string     `db:"name" json:"name"`
	PasswordHash string     `db:"password_hash" json:"-"`
	Role         string     `db:"role" json:"role"` // "client" or "freelancer"
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

// Project represents a job post created by a client
type Project struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ClientID    uuid.UUID  `db:"client_id" json:"client_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	Status      string     `db:"status" json:"status"` // "draft","active","completed","cancelled"
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

// Proposal represents a freelancer’s bid on a project
type Proposal struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	ProjectID    uuid.UUID  `db:"project_id" json:"project_id"`
	FreelancerID uuid.UUID  `db:"freelancer_id" json:"freelancer_id"`
	Status       string     `db:"status" json:"status"` // "submitted","accepted","rejected","withdrawn"
	Content      string     `db:"content" json:"content"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

// Milestone represents a breakdown of a project
type Milestone struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ProjectID   uuid.UUID  `db:"project_id" json:"project_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	DueDate     *time.Time `db:"due_date" json:"due_date,omitempty"`
	Status      string     `db:"status" json:"status"` // "pending","in_progress","completed"
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

// ChatChannel links a project/proposal to a GetStream channel
type ChatChannel struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	ProjectID       *uuid.UUID `db:"project_id" json:"project_id,omitempty"`
	ProposalID      *uuid.UUID `db:"proposal_id" json:"proposal_id,omitempty"`
	StreamChannelID string     `db:"stream_channel_id" json:"stream_channel_id"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
