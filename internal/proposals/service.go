package proposals

import (
	"time"

	"chatapp/v/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

// SubmitProposal creates a new proposal
func (s *Service) SubmitProposal(projectID, freelancerID uuid.UUID, content string) (*models.Proposal, error) {
	proposal := &models.Proposal{
		ID:           uuid.New(),
		ProjectID:    projectID,
		FreelancerID: freelancerID,
		Status:       "submitted",
		Content:      content,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.Repo.Create(proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

// AcceptProposal updates status to accepted
func (s *Service) AcceptProposal(id uuid.UUID) error {
	return s.Repo.UpdateStatus(id, "accepted")
}

// RejectProposal updates status to rejected
func (s *Service) RejectProposal(id uuid.UUID) error {
	return s.Repo.UpdateStatus(id, "rejected")
}
