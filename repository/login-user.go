package repository

import (
	"day-22/model"
	"fmt"
)

func (r *UserRepositoryDB) GetUserLogin(user model.User) (*model.User, error) {
	query := `SELECT id, name, username, password FROM users WHERE username=$1 AND password=$2`

	var userResponse model.User

	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&userResponse.ID, &userResponse.Name, &userResponse.Username, &userResponse.Password)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (r *UserRepositoryDB) UpdateUser(user *model.User) error {
	query := `UPDATE users SET name = $1, username = $2, password = $3, status = $4, token = $5 WHERE id = $6`
	_, err := r.DB.Exec(query, user.Name, user.Username, user.Password, user.Status, user.Token, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}
