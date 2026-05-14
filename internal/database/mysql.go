package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// dsn := "root:password@tcp(localhost:3306)/food_ordering?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
