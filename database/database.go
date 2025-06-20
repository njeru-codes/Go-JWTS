package database

import (
	"log"
	"time"

	model "github.com/njeru-codes/Go-JWTS/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(dsn string) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// run db migrations
	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
	)
	if err != nil{
		log.Fatalf("failed to run db migrations: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)                  // minimum idle connections
	sqlDB.SetMaxOpenConns(100)                 // maximum open connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // maximum connection reuse time

	DB = db
}
