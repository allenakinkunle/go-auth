package auth

import "github.com/golang-jwt/jwt"

type MockTokenManager struct{}

func NewMockTokenManager() MockTokenManager {
	return MockTokenManager{}
}

func (MockTokenManager) NewJWT(subject string) (*Tokens, error) {
	return &Tokens{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil
}

func (m MockTokenManager) ValidateJWT(token string) (*jwt.Token, string, error) {
	return nil, "nanoid", nil
}

func (m MockTokenManager) GenerateRandomToken() string {
	return "randomToken"
}
