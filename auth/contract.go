package auth

type Auth interface {
	GenerateAccessToken()
	GenerateRefreshToken()
	ValidateAccessToken()
	ValidateRefreshToken()
	RotateRefreshToken()
}
