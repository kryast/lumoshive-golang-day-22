package repository

import (
	"database/sql"
	"day-22/model"
	"fmt"
)

func (r *UserRepositoryDB) GetDetailUsers(id int) (*model.User, error) {
	query := `SELECT name, username, password, status, token FROM users WHERE id = $1`
	rows := r.DB.QueryRow(query, id)

	var user model.User
	err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Status, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	// Mengembalikan data pengguna
	return &user, nil
}
