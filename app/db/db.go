package db

import (
	"be-map-test/app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	DB.AutoMigrate(&models.TrackData{})
}
