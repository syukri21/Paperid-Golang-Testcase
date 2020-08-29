package db

import (
	"fmt"
	"os"

	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

var conn *gorm.DB
var err error

// Connection -> create connection with credentials
func Connection() {

	dbAuth := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	conn, err = gorm.Open("mysql", dbAuth)
	if err != nil {
		logging.Error("DB", err)
	}
}

// SeedConnection is method to create connection for seeder
func SeedConnection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// AppConnection -> method to create connection for application
func AppConnection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// TestConnection -> method to create connection for application testing
func TestConnection() {
	if err := godotenv.Load("../.env.test"); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// GetDB -> method to get connection instance
func GetDB() *gorm.DB {
	return conn
}

// DropAllTable -> method to drop all database table (using this only for testing)
func DropAllTable() {
	conn.DropTable(&entity.User{})
}
