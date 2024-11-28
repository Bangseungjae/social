package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TestAuthenticator struct {
}

const secret = "test"

var testClaims = jwt.MapClaims{
	"sub": int64(42),
	"exp": time.Now().Add(time.Hour).Unix(),
	"iat": time.Now().Unix(),
	"nbf": time.Now().Unix(),
	"iss": "test-aud",
	"aud": "test-aud",
}

func (t TestAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, testClaims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString, nil
}

func (t TestAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
