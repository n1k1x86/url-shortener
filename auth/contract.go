package auth

type Auth interface {
	GenerateTokenPair(userID int64, login string) (string, string, error)
	GenerateAccessToken(userID int64, login string) (string, error)
	GenerateRefreshToken(userID int64, login string) (string, error)
	ValidateAccessToken(access string) (bool, error)
	ValidateRefreshToken(refresh string) (bool, error)
	RotateRefreshToken(refresh string) (string, string, error)
}
