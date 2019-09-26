package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PDB struct {
	DB *sql.DB
}

func (pdb PDB) LoadByID(id int) ([]User, error) {
	rows, err := pdb.DB.Query("SELECT * FROM users WHERE id=$1 AND status=TRUE", id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Login, &user.MoneyAmount, &user.CardNumber, &user.Status)
		if err != nil {
			return nil, err
		}
		if user.Status {
			users = append(users, user)
		}
	}
	return users, nil
}

func (pdb PDB) LoadAll() ([]User, error) {
	rows, err := pdb.DB.Query("SELECT * FROM users WHERE status=TRUE")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Login, &user.MoneyAmount, &user.CardNumber, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (pdb PDB) LoadByLogin(login string) ([]User, error) {
	rows, err := pdb.DB.Query("SELECT * FROM users WHERE login=$1 AND status=TRUE", login)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Login, &user.MoneyAmount, &user.CardNumber, &user.Status)
		if err != nil {
			return nil, err
		}
		if user.Status {
			users = append(users, user)
		}
	}
	return users, nil
}
