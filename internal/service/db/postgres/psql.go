package postgres

import (
	"database/sql"
	"fmt"
	"github.com/lll-phill-lll/hsesec/internal/service/db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "bankuser"
	password = "bankpassword"
	dbname   = "bank"
)

type postgresDB struct {
	DB   *sql.DB
}

func New() (db.DataBase, error) {
	postgres := postgresDB{}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	postgres.DB = database

	return postgres, nil
}
