package jwt_token

import (
	"github.com/golang-jwt/jwt/v5"
	"mitra-kirim-be-mgmt/internal/user/model"
	"time"
)

type ConfigJWT struct {
	JwtKey          []byte
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
}

func GenerateAccessToken(userID, email string, jwtKey []byte, exp time.Duration) (string, error) {
	return generateToken(userID, email, "access", exp, jwtKey)
}

func GenerateRefreshToken(userID, email string, jwtKey []byte, exp time.Duration) (string, error) {
	return generateToken(userID, email, "refresh", exp, jwtKey)
}

func generateToken(userID, email, tokenType string, expiration time.Duration, jwtKey []byte) (string, error) {
	claims := &model.Claims{
		UserID:    userID,
		Email:     email,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
