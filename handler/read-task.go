package handler

import (
	"day-22/library"
	"net/http"
)

func (th *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {

	// Call the service to get all tasks
	tasks, err := th.serviceTask.GetAllDataTask()
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	library.SuccessResponse(w, "success get all tasks data", tasks)
}
