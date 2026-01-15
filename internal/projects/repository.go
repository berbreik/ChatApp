package projects

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

// Create inserts a new project
func (r *Repository) Create(project *models.Project) error {
	_, err := r.DB.NamedExec(`
        INSERT INTO projects (id, client_id, title, description, status, created_at, updated_at)
        VALUES (:id, :client_id, :title, :description, :status, :created_at, :updated_at)
    `, project)
	return err
}

// GetByID fetches a project by ID
func (r *Repository) GetByID(id uuid.UUID) (*models.Project, error) {
	var project models.Project
	err := r.DB.Get(&project, `SELECT * FROM projects WHERE id=$1 AND deleted_at IS NULL`, id)
	return &project, err
}

// ListByClient returns all active projects for a client
func (r *Repository) ListByClient(clientID uuid.UUID) ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.Select(&projects, `SELECT * FROM projects WHERE client_id=$1 AND deleted_at IS NULL`, clientID)
	return projects, err
}

// SoftDelete marks a project as deleted
func (r *Repository) SoftDelete(id uuid.UUID) error {
	_, err := r.DB.Exec(`UPDATE projects SET deleted_at = now() WHERE id=$1`, id)
	return err
}
