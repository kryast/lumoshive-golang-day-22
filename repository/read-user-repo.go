package repository

import "day-22/model"

func (r *UserRepositoryDB) GetAll() (*[]model.User, error) {
	users := []model.User{}
	query := `SELECT id, name, username, password, status, token FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Status, &user.Token)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
