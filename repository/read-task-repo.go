package repository

import "day-22/model"

func (r *TaskRepositoryDB) GetAll() (*[]model.Task, error) {
	query := `SELECT description, status FROM tasks`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var task model.Task
		rows.Scan(&task.Description, &task.Status)

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}
