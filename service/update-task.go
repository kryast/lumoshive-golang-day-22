package service

import "fmt"

func (ts *TaskService) UpdateTaskStatus(id int) error {
	// Panggil repository untuk memperbarui status task
	err := ts.RepoTask.UpdateStatusTask(id)
	if err != nil {
		return fmt.Errorf("failed to update task status: %v", err)
	}
	return nil
}
