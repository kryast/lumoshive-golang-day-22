package handler

import (
	"day-22/library"
	"day-22/model"
	"day-22/service"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	serviceTask service.TaskService
}

func NewTaskHandler(ts service.TaskService) TaskHandler {
	return TaskHandler{serviceTask: ts}
}

func (th *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	var tasks map[string]interface{}
	tasks, err = th.serviceTask.InputDataTask(task.Description)
	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	library.SuccessResponse(w, "success create task", tasks)
}
