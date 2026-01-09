package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user account
type User struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"-"` // Never expose password hash in JSON
	FullName     string     `db:"full_name" json:"full_name"`
	Role         string     `db:"role" json:"role"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	LastLoginAt  *time.Time `db:"last_login_at" json:"last_login_at,omitempty"`
}

// Session represents a refresh token session
type Session struct {
	ID           uuid.UUID `db:"id" json:"id"`
	UserID       uuid.UUID `db:"user_id" json:"user_id"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	UserAgent    string    `db:"user_agent" json:"user_agent"`
	IPAddress    string    `db:"ip_address" json:"ip_address"`
	ExpiresAt    time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
