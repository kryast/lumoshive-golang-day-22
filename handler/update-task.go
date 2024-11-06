package handler

import (
	"net/http"
	"strconv"
)

func (th *TaskHandler) UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = th.serviceTask.UpdateTaskStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks, err := th.serviceTask.GetAllDataTask() // Pastikan ada method untuk ambil semua task
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	// Render template untuk menampilkan daftar task
	templates.ExecuteTemplate(w, "list-task-view", tasks)
}
