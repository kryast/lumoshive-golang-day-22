package handler

import (
	"day-22/library"
	"day-22/model"
	"day-22/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	serviceUsers service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return UserHandler{serviceUsers: us}
}

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	var users map[string]interface{}
	users, err = uh.serviceUsers.InputDataUser(user.Name, user.Username, user.Password)
	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	library.SuccessResponse(w, "Success create user", users)
	http.Redirect(w, r, "/all-customer", http.StatusSeeOther)
}
