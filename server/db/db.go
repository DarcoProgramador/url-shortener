package db

import (
	"log"

	"github.com/DarcoProgramador/url-shortener/config"
	"github.com/DarcoProgramador/url-shortener/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(conf *config.Config) *gorm.DB {
	dbURL := conf.DatabaseURL

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Url{})

	return db
}
