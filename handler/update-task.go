package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (th *TaskHandler) UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID task dari URL parameter (misalnya /task/{id}/status)
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Menggunakan service untuk memperbarui status task

	err = th.serviceTask.UpdateTaskStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"StatusCode": 200,
		"Message":    "Task status updated successfully",
	})
}
