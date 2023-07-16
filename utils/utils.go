package utils

import (
	"time"

	"github.com/go-redis/redis"
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

func InitRedisCacheTest() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
