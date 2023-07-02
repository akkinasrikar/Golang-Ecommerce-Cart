package initializers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "ecom"
)

var (
	Db  *sql.DB
	err error
)

func ConnectDataBase() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	Db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	err = Db.Ping()
	CheckError(err)
	fmt.Println("Successfully connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
