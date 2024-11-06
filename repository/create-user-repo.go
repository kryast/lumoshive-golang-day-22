package repository

import (
	"database/sql"
	"day-22/model"
	"errors"
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

func (repo *UserRepositoryDB) GetUserByToken(token string) (model.User, error) {
	// Query untuk mencari user berdasarkan token
	var user model.User
	query := "SELECT id, name, username, password, status, token FROM users WHERE token = ?"
	err := repo.DB.QueryRow(query, token).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Status, &user.Token)

	// Jika token tidak ditemukan
	if err != nil {
		if err == sql.ErrNoRows {
			// Token tidak ditemukan
			return user, errors.New("token tidak ditemukan")
		}
		// Error lainnya
		return user, fmt.Errorf("error querying database: %v", err)
	}

	// Kembalikan user yang ditemukan
	return user, nil
}
