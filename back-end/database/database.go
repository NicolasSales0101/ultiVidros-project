package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_ROOT_PASSWORD := os.Getenv("MYSQL_ROOT_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_ROOT_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error in connection to database:", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
