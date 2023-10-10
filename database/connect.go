package database

import (
	"log"
	"os"

	"netflix/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&models.Library{}, &models.Collection{}, &models.Video{})
	if err != nil {
		log.Fatal(err)
	}

}
