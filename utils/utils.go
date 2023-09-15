package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(Username string, UserId int64) (string, error) {
	claims := jwt.MapClaims{
		"sub":     Username,
		"usersId": UserId,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
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

func SetContextWithAuthData() context.Context {
	var authData models.AuthData
	authData.UsersId = int64(1234)
	ctx := context.Background()
	ctx = context.WithValue(ctx, models.EcomctxKey("AuthData"), authData)
	return ctx
}

func GenerateRandomUserIdNumber() int {
	return 100000 + rand.Intn(899999)
}

func GenerateRandomString() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateEcomId() string {
	return "ecom_" + GenerateRandomString()
}

func GenerateCardId() string {
	return "card_" + GenerateRandomString()
}

func ValidateUnkownParams(ctx *gin.Context, body interface{}) models.EcomError {
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrorMessage.JSON_UNKNOWN_FIELD) {
			param := strings.TrimLeft(err.Error(), constants.ErrorMessage.JSON_UNKNOWN_FIELD)
			param = strings.Replace(param, "\"", "", -1)
			return *helper.ErrorUnknownParam(param)
		} else if strings.Contains(err.Error(), constants.ErrorMessage.JSON_CANNOT_UNMARSHAL) {
			param := err.Error()[strings.LastIndex(err.Error(), ".")+1:]
			if strings.Contains(param, " ") {
				param = param[:strings.Index(param, " ")]
			}
			expectedDataType := err.Error()[strings.LastIndex(err.Error(), " ")+1:]
			return *helper.ErrorParamMissingOrInvalid(param, expectedDataType)
		}
	}

	payloadBS, err := json.Marshal(&body)
	if err != nil {
		return *helper.ErrorParamMissingOrInvalid(err.Error(), "payload")
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(payloadBS))
	return models.EcomError{}
}

func ValidateCardExpiryDate(expiryDate string) bool {
	layout := "01/06"
	t, err := time.Parse(layout, expiryDate)
	if err != nil {
		return false
	}
	if t.Before(time.Now()) {
		return false
	}
	return true
}
