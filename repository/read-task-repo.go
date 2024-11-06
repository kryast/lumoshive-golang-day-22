package repository

import "day-22/model"

func (r *TaskRepositoryDB) GetAll() (*[]model.Task, error) {
	tasks := []model.Task{}
	query := `SELECT id, description, status FROM tasks`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task model.Task
		rows.Scan(&task.ID, &task.Description, &task.Status)

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}
