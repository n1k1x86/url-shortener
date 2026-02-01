package auth

import "context"

type Auth interface {
	GenerateTokenPair(ctx context.Context, userID int64, login string) (string, string, string, error)
	GenerateAccessToken(ctx context.Context, userID int64, login string) (string, error)
	GenerateRefreshToken(ctx context.Context, userID int64, login string) (string, string, error)
	ValidateAccessToken(ctx context.Context, access string) (int64, bool, error)
	RefreshAccessToken(ctx context.Context, refresh string) (string, string, error)
}
