package service

import (
	"fmt"
)

func (ts *TaskService) GetAllDataTask() ([]map[string]interface{}, error) {

	tasks, err := ts.RepoTask.GetAll()

	if err != nil {
		return nil, err
	}

	if tasks == nil || len(*tasks) == 0 {
		fmt.Println("No task found")
		return nil, err
	}

	var listTask []map[string]interface{}
	for _, task := range *tasks {
		listTask = append(listTask, map[string]interface{}{
			"description": task.Description,
			"status":      task.Status,
		})
	}

	return listTask, nil

}
