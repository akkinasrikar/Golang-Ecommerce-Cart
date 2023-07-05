package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(Username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": Username,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	signingMethod := jwt.SigningMethodHS256
	secretKey := []byte("testing")

	token := jwt.NewWithClaims(signingMethod, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
