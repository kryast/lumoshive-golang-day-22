package handler

import (
	"day-22/database"
	"day-22/repository"
	"day-22/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID task dari URL parameter (misalnya /task/{id}/status)
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Menggunakan service untuk memperbarui status task
	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo) // Gantilah dengan repository yang sesuai
	err = taskService.UpdateTaskStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"StatusCode": 200,
		"Message":    "Task status updated successfully",
	})
}
