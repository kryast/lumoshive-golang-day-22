package handler

import (
	"day-22/database"
	"day-22/model"
	"day-22/repository"
	"day-22/service"
	"encoding/json"
	"net/http"
)

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create the repository and service
	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)

	// Call the service to get all tasks
	tasks, err := taskService.GetAllDataTask()
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "user retrieved successfully",
		Data:       tasks,
	}

	json.NewEncoder(w).Encode(response)
}
