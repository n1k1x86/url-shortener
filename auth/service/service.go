package service

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"
	auth_config "url-shortener/auth/config"
	auth_models "url-shortener/auth/models"
	"url-shortener/auth/repo"

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
	cfg  *auth_config.Config
	repo *repo.Repository
}

func NewService(cfg *auth_config.Config, repo *repo.Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) GenerateTokenPair(ctx context.Context, userID int64, login string) (string, string, string, error) {
	access, err := s.GenerateAccessToken(ctx, userID, login)
	if err != nil {
		return "", "", "", err
	}
	refresh, jti, err := s.GenerateRefreshToken(ctx, userID, login)
	if err != nil {
		return "", "", "", err
	}
	return access, refresh, jti, nil
}

func (s *Service) GenerateAccessToken(ctx context.Context, userID int64, login string) (string, error) {
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

func (s *Service) GenerateRefreshToken(ctx context.Context, userID int64, login string) (string, string, error) {
	jti := uuid.New().String()
	refreshClaims := &auth_models.CustomClaims{
		UserID: userID,
		Login:  login,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.cfg.Base.Issuer,
			Audience:  jwt.ClaimStrings{s.cfg.Base.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.RefreshToken.ExpiredAfter)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jti,
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err := refreshToken.SignedString([]byte(s.cfg.RefreshToken.Secret))
	if err != nil {
		return "", "", err
	}

	err = s.repo.InsertRefreshToken(ctx, s.getRefreshHash(refresh), refreshClaims.ID, userID)

	if err != nil {
		return "", "", err
	}

	return refresh, jti, err
}

func (s *Service) ValidateAccessToken(ctx context.Context, access string) (int64, bool, error) {
	claims := &auth_models.CustomClaims{}
	token, err := jwt.ParseWithClaims(access, claims, func(t *jwt.Token) (any, error) {
		return []byte(s.cfg.AccessToken.Secret), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithIssuer(s.cfg.Base.Issuer),
		jwt.WithExpirationRequired(),
	)

	if err != nil || !token.Valid {
		return 0, false, err
	}

	if ok, err := s.repo.IsUserExist(ctx, claims.Login, claims.UserID); !ok || err != nil {
		if err != nil {
			return 0, false, err
		}
		return 0, false, fmt.Errorf("user does not exist")
	}

	return claims.UserID, true, nil
}

func (s *Service) RefreshAccessToken(ctx context.Context, refresh string) (string, string, error) {
	claims := &auth_models.CustomClaims{}
	token, err := jwt.ParseWithClaims(refresh, claims, func(t *jwt.Token) (any, error) {
		return []byte(s.cfg.RefreshToken.Secret), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithIssuer(s.cfg.Base.Issuer),
		jwt.WithExpirationRequired(),
	)

	if err != nil || !token.Valid {
		return "", "", err
	}

	if ok, err := s.repo.IsTokenRevoked(ctx, claims.ID); ok || err != nil {
		if err != nil {
			return "", "", err
		}
		return "", "", fmt.Errorf("token is already revoked")
	}

	access, refresh, jti, err := s.GenerateTokenPair(ctx, claims.UserID, claims.Login)
	if err != nil {
		return "", "", err
	}

	err = s.repo.RevokeToken(ctx, jti, claims.ID)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *Service) getRefreshHash(refresh string) string {
	hashSum := sha512.Sum512([]byte(refresh))
	return hex.EncodeToString(hashSum[:])
}
