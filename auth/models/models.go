package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshToken struct {
	ID         int64
	JTI        string
	Hash       string
	UserID     int64
	ReplacedBy string
	IsRevoked  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CustomClaims struct {
	UserID int64
	Login  string
	jwt.RegisteredClaims
}
