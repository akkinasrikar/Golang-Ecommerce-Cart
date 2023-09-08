package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/akkinasrikar/ecommerce-cart/api/dto"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "ecom"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDataBase() *gorm.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	Db, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	CheckError(err)

	sqlDB, err := Db.DB()
	CheckError(err)

	err = sqlDB.Ping()
	CheckError(err)

	err = AutoMigrate(Db)
	CheckError(err)

	fmt.Println("Successfully connected and AutoMigrated!")
	return Db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func AutoMigrate(db *gorm.DB) error {
	// automigrate with table name
	err := db.AutoMigrate(&entities.SignUp{})
	if err != nil {
		return err
	}
	return nil
}

func IsDataSeeded(dbStore *db) bool {
	count := dbStore.Count(&entities.Item{})
	return count > 0
}

func SeedData(database *db) {
 	var items []entities.Item
	var itemsDto []dto.Items

	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://fakestoreapi.com/products", nil)
	if err != nil {
		panic(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(responseBody, &itemsDto)
	if err != nil {
		panic(err)
	}
	for _, item := range itemsDto {
		items = append(items, entities.Item{
			ItemID:          item.Id,
			ItemTitle:       item.Title,
			ItemPrice:       item.Price,
			ItemDescription: item.Description,
			ItemCategory:    item.Category,
			ItemImage:       item.Image,
			ItemRating:      item.Rating.Rate,
			ItemCount:       item.Rating.Count,
		})
	}
	database.Create(&items)
}
