package db

import (
	"database/sql"
	"fmt"
	"github.com/lll-phill-lll/hsesec/internal/service/db/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "bankuser"
	password = "bankpassword"
	dbname   = "bank"
)

type DataBase interface {
	LoadByID(id int) ([]postgres.User, error)
	LoadAll() ([]postgres.User, error)
	LoadByLogin(login string) ([]postgres.User, error)
}

func New() (DataBase, error) {
	p := postgres.PDB{}

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
	p.DB = database

	return p, nil
}
