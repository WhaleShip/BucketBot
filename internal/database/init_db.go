package database

import (
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host                 string
	Port                 string
	Username             string
	Password             string
	DBName               string
	SSLMode              string
	PreferSimpleProtocol bool
}

func getEnvVariable(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalf("Enviroment error: %s variable not exist", name)
	}
	return value
}

func GetInitializedDb() (*pgx.Conn, error) {
	conn, err := ConnectPostgres(Config{
		Host:                 getEnvVariable("PGBOUNCER_HOST"),
		Port:                 getEnvVariable("PGBOUNCER_PORT"),
		Username:             getEnvVariable("POSTGRES_USER"),
		Password:             getEnvVariable("POSTGRES_PASSWORD"),
		DBName:               getEnvVariable("POSTGRES_DB"),
		SSLMode:              getEnvVariable("SSL_MODE"),
		PreferSimpleProtocol: true,
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
