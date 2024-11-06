package handler

import (
	"net/http"
)

func (th *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := th.serviceTask.GetAllDataTask()
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// library.SuccessResponse(w, "success get all tasks data", tasks)
	templates.ExecuteTemplate(w, "list-task-view", tasks)
}
