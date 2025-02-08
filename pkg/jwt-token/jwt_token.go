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

func GenerateAccessToken(userID, username string, jwtKey []byte, exp time.Duration) (string, error) {
	return generateToken(userID, username, "access", exp, jwtKey)
}

func GenerateRefreshToken(userID, username string, jwtKey []byte, exp time.Duration) (string, error) {
	return generateToken(userID, username, "refresh", exp, jwtKey)
}

func generateToken(userID, username, tokenType string, expiration time.Duration, jwtKey []byte) (string, error) {
	claims := &model.Claims{
		UserID:    userID,
		Username:  username,
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
