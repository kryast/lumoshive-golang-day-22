package handler

import (
	"day-22/library"
	"net/http"
)

func (uh *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	// Call the service to get all tasks
	users, err := uh.serviceUsers.GetAllDataUser()
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	library.SuccessResponse(w, "success get all data users", users)
}
