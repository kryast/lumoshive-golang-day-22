package handler

import (
	"day-22/library"
	"net/http"
	"strconv"
)

func (uh *UserHandler) GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uh.serviceUsers.GetUserDetails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	library.SuccessResponse(w, "success get users detail", user)
}
