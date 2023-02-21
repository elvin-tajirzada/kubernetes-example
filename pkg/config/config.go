package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ADDR string
	DB   DB
}

type DB struct {
	Postgres Postgres
}
type Postgres struct {
	Host, Port, User, Password, DBName, SSLMode string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}
	return &Config{
		ADDR: os.Getenv("ADDR"),
		DB: DB{
			Postgres: Postgres{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DBName:   os.Getenv("POSTGRES_DB_NAME"),
				SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
			},
		},
	}, nil
}
