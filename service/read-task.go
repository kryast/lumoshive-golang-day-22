package service

import (
	"day-22/model"
)

func (ts *TaskService) GetAllDataTask() (*[]model.Task, error) {

	tasks, err := ts.RepoTask.GetAll()

	if err != nil {
		return nil, err
	}

	return tasks, nil

}
