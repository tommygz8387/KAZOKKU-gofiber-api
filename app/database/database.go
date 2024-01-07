package database

import (
	"log"
	"os"
	"v1/app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

    // Construct the database DSN
    dsn := Config("DB_USER") + ":" + Config("DB_PASSWORD") + "@tcp(" + Config("DB_HOST") + ":3306)/" + Config("DB_NAME") + "?parseTime=true"

    // Initialize database connection
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
    if err != nil {
        log.Fatalln("could not connect to database")
    }

	// Running Log
    log.Println("Connected")

	// Auto Migrate
	if os.Getenv("DB_AUTO_MIGRATE") == "true" {
		log.Println("Running AutoMigrations")
		db.AutoMigrate(&models.User{},&models.UserPhoto{},&models.UserCreditCard{})
	}

	// Store the database connection globally
	DB = db

    return db, nil
}

// Config returns the database configuration from the .env file
func Config(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
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