package users

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

// Create inserts a new user
func (r *Repository) Create(user *models.User) error {
	_, err := r.DB.NamedExec(`
        INSERT INTO users (id, email, name, password_hash, role, created_at, updated_at)
        VALUES (:id, :email, :name, :password_hash, :role, :created_at, :updated_at)
    `, user)
	return err
}

// GetByEmail fetches a user by email
func (r *Repository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Get(&user, `SELECT * FROM users WHERE email=$1 AND deleted_at IS NULL`, email)
	return &user, err
}

// GetByID fetches a user by ID
func (r *Repository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.DB.Get(&user, `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`, id)
	return &user, err
}

// List returns all active users
func (r *Repository) List() ([]models.User, error) {
	var users []models.User
	err := r.DB.Select(&users, `SELECT * FROM users WHERE deleted_at IS NULL`)
	return users, err
}

// SoftDelete marks a user as deleted
func (r *Repository) SoftDelete(id uuid.UUID) error {
	_, err := r.DB.Exec(`UPDATE users SET deleted_at = now() WHERE id=$1`, id)
	return err
}
