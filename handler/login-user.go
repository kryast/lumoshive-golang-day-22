package handler

import (
	"day-22/library"
	"day-22/model"
	"net/http"
)

// UserLoginHandler processes login attempts from the login form
func (uh *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	// Create user object from parsed form
	user := model.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// Call the service to attempt login
	userData, err := uh.serviceUsers.LoginUser(user)
	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	// On success, render the success view
	templates.ExecuteTemplate(w, "login-success-view", userData)
}
