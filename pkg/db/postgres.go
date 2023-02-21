package db

import (
	"fmt"
	"github.com/elvin-tacirzade/kubernetes-example/pkg/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type DB struct {
	Postgres *sqlx.DB
}

func Connect(conf *config.Config) (*DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.DB.Postgres.Host, conf.DB.Postgres.Port, conf.DB.Postgres.User, conf.DB.Postgres.Password, conf.DB.Postgres.DBName, conf.DB.Postgres.SSLMode)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres: %v", err)
	}

	return &DB{
		Postgres: db,
	}, err
}
