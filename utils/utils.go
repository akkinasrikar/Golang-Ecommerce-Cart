package utils

import (
	"math/rand"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(Username string, UserId int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": Username,
		"usersId": UserId,
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

func SetContext() *gin.Context {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	return ctx
}

func GenerateRandomUserIdNumber() int {
	return 100000 + rand.Intn(899999)
}
