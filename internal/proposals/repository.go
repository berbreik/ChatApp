package proposals

import (
	"chatapp/v/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

// Create inserts a new proposal
func (r *Repository) Create(p *models.Proposal) error {
	_, err := r.DB.NamedExec(`
        INSERT INTO proposals (id, project_id, freelancer_id, status, content, created_at, updated_at)
        VALUES (:id, :project_id, :freelancer_id, :status, :content, :created_at, :updated_at)
    `, p)
	return err
}

// GetByID fetches a proposal by ID
func (r *Repository) GetByID(id uuid.UUID) (*models.Proposal, error) {
	var proposal models.Proposal
	err := r.DB.Get(&proposal, `SELECT * FROM proposals WHERE id=$1 AND deleted_at IS NULL`, id)
	return &proposal, err
}

// ListByProject returns all proposals for a project
func (r *Repository) ListByProject(projectID uuid.UUID) ([]models.Proposal, error) {
	var proposals []models.Proposal
	err := r.DB.Select(&proposals, `SELECT * FROM proposals WHERE project_id=$1 AND deleted_at IS NULL`, projectID)
	return proposals, err
}

// UpdateStatus changes the status of a proposal
func (r *Repository) UpdateStatus(id uuid.UUID, status string) error {
	_, err := r.DB.Exec(`UPDATE proposals SET status=$1, updated_at=now() WHERE id=$2 AND deleted_at IS NULL`, status, id)
	return err
}
