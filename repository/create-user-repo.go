package repository

import (
	"database/sql"
	"day-22/model"
)

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{DB: db}
}

func (r *UserRepositoryDB) CreateDataUser(user model.User) error {
	query := "INSERT INTO users (name , username, password, status) VALUES ($1 , $2 , $3, 'InActive') RETURNING id"

	err := r.DB.QueryRow(query, user.Name, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
