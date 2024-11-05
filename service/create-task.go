package service

import (
	"day-22/model"
	"day-22/repository"
	"errors"
	"fmt"
)

type TaskService struct {
	RepoTask repository.TaskRepositoryDB
}

func NewTaskService(repo repository.TaskRepositoryDB) *TaskService {
	return &TaskService{RepoTask: repo}
}

// InputDataTask untuk memasukkan task baru dan mengembalikan data yang berisi description dan status
func (ts *TaskService) InputDataTask(description string) (map[string]interface{}, error) {
	if description == "" {
		return nil, errors.New("description tidak boleh kosong")
	}

	task := model.Task{
		Description: description,
		Status:      "incomplete",
	}

	// Menyimpan task ke dalam database
	err := ts.RepoTask.CreateDataTask(task)
	if err != nil {
		fmt.Println("Error :", err)
		return nil, err
	}

	// Membuat map untuk menampilkan description dan status saja
	response := map[string]interface{}{
		"Description": task.Description,
		"Status":      task.Status,
	}

	// Menampilkan pesan sukses
	fmt.Println("Berhasil input data task dengan description:", task.Description)

	// Mengembalikan map yang berisi data task (Description dan Status)
	return response, nil
}
