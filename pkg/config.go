package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
  envErr := godotenv.Load(".env")
  if envErr != nil {
    panic(envErr)
  }
	d, err := gorm.Open("mysql", os.Getenv("DB_DNS")) // Use your local DB
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
