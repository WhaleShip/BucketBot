package database

import (
	"log"
	"os"

	"github.com/WhaleShip/BucketBot/internal/database/models"
	"gorm.io/gorm"
)

func getEnvVariable(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalf("Enviroment error: %s variable not exist", name)
	}
	return value
}

func GetInitializedDb() (*gorm.DB, error) {
	db, err := ConnectPostgres(Config{
		Host:     getEnvVariable("POSTGRES_HOST"),
		Port:     getEnvVariable("POSTGRES_PORT"),
		Username: getEnvVariable("POSTGRES_USER"),
		Password: getEnvVariable("POSTGRES_PASSWORD"),
		DBName:   getEnvVariable("POSTGRES_DB"),
		SSLMode:  getEnvVariable("SSL_MODE"),
	})
	if err != nil {
		log.Fatal("db connection fail: ", err.Error())
		return nil, err
	}

	err = db.AutoMigrate(&models.Note{}, &models.UserNotes{})
	if err != nil {
		log.Fatal("automigrate error: ", err.Error())
		return db, err
	}

	return db, nil
}
