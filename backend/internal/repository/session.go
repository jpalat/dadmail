package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jay/dadmail/internal/models"
	"github.com/jmoiron/sqlx"
)

// SessionRepository handles session database operations
type SessionRepository struct {
	db *sqlx.DB
}

// NewSessionRepository creates a new session repository
func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// Create creates a new session
func (r *SessionRepository) Create(userID uuid.UUID, refreshToken, userAgent, ipAddress string, expiresAt time.Time) (*models.Session, error) {
	session := &models.Session{
		ID:           uuid.New(),
		UserID:       userID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		IPAddress:    ipAddress,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}

	query := `
		INSERT INTO sessions (id, user_id, refresh_token, user_agent, ip_address, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, session.ID, session.UserID, session.RefreshToken, session.UserAgent, session.IPAddress, session.ExpiresAt, session.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

// GetByRefreshToken retrieves a session by refresh token
func (r *SessionRepository) GetByRefreshToken(refreshToken string) (*models.Session, error) {
	session := &models.Session{}
	query := `SELECT * FROM sessions WHERE refresh_token = $1 AND expires_at > NOW()`

	err := r.db.Get(session, query, refreshToken)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("session not found or expired")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	return session, nil
}

// Delete deletes a session by refresh token
func (r *SessionRepository) Delete(refreshToken string) error {
	query := `DELETE FROM sessions WHERE refresh_token = $1`

	_, err := r.db.Exec(query, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

// DeleteAllForUser deletes all sessions for a user
func (r *SessionRepository) DeleteAllForUser(userID uuid.UUID) error {
	query := `DELETE FROM sessions WHERE user_id = $1`

	_, err := r.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user sessions: %w", err)
	}

	return nil
}

// DeleteExpired deletes all expired sessions
func (r *SessionRepository) DeleteExpired() error {
	query := `DELETE FROM sessions WHERE expires_at <= NOW()`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to delete expired sessions: %w", err)
	}

	return nil
}
