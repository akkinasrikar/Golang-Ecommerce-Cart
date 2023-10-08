package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nfnt/resize"
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

func GenerateAddressId() string {
	return "address_" + GenerateRandomString()
}

func GenerateOrderId() string {
	return "order_" + GenerateRandomString()
}

func GenerateTaskID() string {
	return "task_" + GenerateRandomString()
}

func GenerateRandomDate() string {
	min := time.Now().Unix()
	max := time.Now().AddDate(0, 0, 3).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}

func GenerateCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func GetDeliveryAddress(data entities.DeliveryAddress) string {
	return data.HouseNo + ", " + data.Street + ", " + data.City + ", " + data.State + ", " + data.Pincode
}

func FormatCardNumber(cardNumber int64) string {
	cardNumberStr := strconv.FormatInt(cardNumber, 10)
	return "XXXX-XXXX-XXXX-" + cardNumberStr[len(cardNumberStr)-4:]
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

func UnmarshallCartItems(data string) (entities.ItemsInCart, error) {
	var cartItems entities.ItemsInCart
	err := json.Unmarshal([]byte(data), &cartItems)
	if err != nil {
		return entities.ItemsInCart{}, err
	}
	return cartItems, nil
}

func ReadImageFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	image, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func GenerateHtmlResponse(image string, data entities.Item) string {
	htmlResponse := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Data with Image</title>
		</head>
			<body>
				<h1 style="text-align:center">Product Details</h1>
					<p><strong>Item ID:</strong> %d</p>
					<p><strong>Item Title:</strong> %s</p>
					<p><strong>Item Price:</strong> $%.2f</p>
					<p><strong>Item Description:</strong> %s</p>
					<p><strong>Item Category:</strong> %s</p>
					<p><strong>Item Rating:</strong> %.2f</p>
					<p><strong>Item Count:</strong> %d</p>
				<h1 style="text-align:center" >Image</h1>
					<div style="text-align: center;">
						<img src="data:image/png;base64,%s" alt="Embedded Image" 
					</div>
			</body>
	</html>
`
	htmlResponse = fmt.Sprintf(htmlResponse, data.ItemID, data.ItemTitle, data.ItemPrice, data.ItemDescription, data.ItemCategory, data.ItemRating, data.ItemCount, image)
	return htmlResponse
}

func GenerateHtmlResponse2(data entities.Item, orderDetails entities.Order) string {
	htmlResponse := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Data with Image</title>
		</head>
			<body>
				<h1 style="text-align:center">Product Details</h1>
					<p><strong>Item ID:</strong> %d</p>
					<p><strong>Item Title:</strong> %s</p>
					<p><strong>Item Price:</strong> $%.2f</p>
					<p><strong>Item Description:</strong> %s</p>
					<p><strong>Item Category:</strong> %s</p>
					<p><strong>Item Rating:</strong> %.2f</p>
					<p><strong>Item Count:</strong> %d</p>
					<p><strong>Order ID:</strong> %s</p>
					<p><strong>Order Date:</strong> %s</p>
					<p><strong>Order Status:</strong> %s</p>
					<p><strong>Order Total:</strong> $%.2f</p>
					<p><strong>Order Delivery Date:</strong> %s</p>
					<p>Thanks for shopping with us!</p>
				<h1 style="text-align:center" >Image</h1>
					<div style="text-align: center;">
						<img src="data:image/png;base64,%s" alt="Embedded Image"
					</div>
			</body>
	</html>
`
	htmlResponse = fmt.Sprintf(htmlResponse, data.ItemID, data.ItemTitle, data.ItemPrice, data.ItemDescription, data.ItemCategory, data.ItemRating, data.ItemCount, orderDetails.OrderID, orderDetails.OrderDate, orderDetails.OrderStatus, orderDetails.OrderAmount, orderDetails.DeliveryDate, data.ImageBase64)
	return htmlResponse
}

func ResizeImage(imageBytes []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}
	newWidth := 200
	newHeight := 200
	resizedImg := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resizedImg, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GeneratePdf(html string) ([]byte, error) {
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	pdf.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	err = pdf.Create()
	if err != nil {
		return nil, err
	}
	return pdf.Bytes(), nil
}
