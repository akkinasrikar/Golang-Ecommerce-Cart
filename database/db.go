package initializers

import (
	"fmt"

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

	fmt.Println("Successfully connected and AutoMigrated!")
	return Db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
