package db

import (
	"log"

	"github.com/asrma7/playpal/feed-svc/config"
	"github.com/asrma7/playpal/feed-svc/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func Init(cfg config.Config) DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	db.AutoMigrate(&models.Feed{})

	return DB{DB: db}
}
