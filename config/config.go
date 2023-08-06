package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Init() {
	_, b, _, _ := runtime.Caller(0)
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(ProjectRootPath + "/.env")
	if err != nil {
		log.Println("Error loading .env file :: " + err.Error())
	}

	loadFakeStoreConfig()
}
