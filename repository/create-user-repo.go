package repository

import (
	"database/sql"
	"day-22/model"
	"fmt"
)

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{DB: db}
}

func (r *UserRepositoryDB) CreateDataUser(user model.User) error {
	query := "INSERT INTO users (name , username, password, status, token) VALUES ($1 , $2 , $3, 'InActive', 'nil') RETURNING id"

	err := r.DB.QueryRow(query, user.Name, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepositoryDB) CheckToken(token string) string {
	var tokenResult string
	query := `SELECT token FROM users WHERE token=$1`
	err := ur.DB.QueryRow(query, token).Scan(&tokenResult)

	fmt.Println(err)
	return tokenResult
}
