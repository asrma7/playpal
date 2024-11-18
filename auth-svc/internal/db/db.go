package db

import (
	"github.com/asrma7/playpal/auth-svc/config"
	"github.com/asrma7/playpal/auth-svc/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	DB *gorm.DB
}

func Init(cfg config.Config) DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	db.AutoMigrate(&models.User{})

	return DB{DB: db}
}
