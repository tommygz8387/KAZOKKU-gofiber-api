package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
    // Load environment variables from .env
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    // Get database configuration from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    // Construct the database DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

    // Initialize database connection
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
    if err != nil {
        panic(err)
    }

	// Running Log
    log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	// Auto Migrate
	// log.Println("Running AutoMigrations")
	// db.AutoMigrate(&models.User{},&models.UserPhoto{},&models.UserCreditCard{})

	// Store the database connection globally
	DB = db

    return db, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return nil
}