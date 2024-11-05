package repository

import "day-22/model"

func (r *UserRepositoryDB) GetAll() (*[]model.User, error) {
	query := `SELECT name, status FROM users`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var user model.User
		rows.Scan(&user.Name, &user.Status)

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
