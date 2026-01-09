package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jay/dadmail/internal/models"
	"github.com/jmoiron/sqlx"
)

// UserRepository handles user database operations
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(email, passwordHash, fullName, role string) (*models.User, error) {
	user := &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		FullName:     fullName,
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `
		INSERT INTO users (id, email, password_hash, full_name, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, user.ID, user.Email, user.PasswordHash, user.FullName, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT * FROM users WHERE email = $1`

	err := r.db.Get(user, query, email)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	query := `SELECT * FROM users WHERE id = $1`

	err := r.db.Get(user, query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// UpdateLastLogin updates the user's last login timestamp
func (r *UserRepository) UpdateLastLogin(id uuid.UUID) error {
	query := `UPDATE users SET last_login_at = $1 WHERE id = $2`
	now := time.Now()

	_, err := r.db.Exec(query, now, id)
	if err != nil {
		return fmt.Errorf("failed to update last login: %w", err)
	}

	return nil
}

// Update updates a user's profile
func (r *UserRepository) Update(id uuid.UUID, fullName string) error {
	query := `UPDATE users SET full_name = $1, updated_at = $2 WHERE id = $3`

	_, err := r.db.Exec(query, fullName, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
