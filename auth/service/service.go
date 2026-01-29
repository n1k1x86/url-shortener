package service

import (
	"time"
	auth_config "url-shortener/auth/config"
	auth_models "url-shortener/auth/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// type Auth interface {
// 	GenerateAccessToken()
// 	GenerateRefreshToken()
// 	ValidateAccessToken()
// 	ValidateRefreshToken()
// 	RotateRefreshToken()
// }

type Service struct {
	cfg *auth_config.Config
}

func NewService(cfg *auth_config.Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) GenerateTokenPair(userID int64, login string) (string, string, error) {
	access, err := s.GenerateAccessToken(userID, login)
	if err != nil {
		return "", "", err
	}
	refresh, err := s.GenerateRefreshToken(userID, login)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (s *Service) GenerateAccessToken(userID int64, login string) (string, error) {
	accessClaims := &auth_models.CustomClaims{
		UserID: userID,
		Login:  login,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.cfg.Base.Issuer,
			Audience:  jwt.ClaimStrings{s.cfg.Base.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.AccessToken.ExpiredAfter)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	access, err := accessToken.SignedString([]byte(s.cfg.AccessToken.Secret))
	if err != nil {
		return "", err
	}
	return access, err
}

func (s *Service) GenerateRefreshToken(userID int64, login string) (string, error) {

	accessClaims := &auth_models.CustomClaims{
		UserID: userID,
		Login:  login,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.cfg.Base.Issuer,
			Audience:  jwt.ClaimStrings{s.cfg.Base.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.RefreshToken.ExpiredAfter)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refresh, err := refreshToken.SignedString([]byte(s.cfg.AccessToken.Secret))
	if err != nil {
		return "", err
	}
	return refresh, err
}
