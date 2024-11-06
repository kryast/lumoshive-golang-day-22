package handler

import (
	"day-22/library"
	"day-22/model"
	"day-22/service"
	"net/http"
)

type UserHandler struct {
	serviceUsers service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return UserHandler{serviceUsers: us}
}

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Jika menggunakan form-urlencoded, gunakan r.FormValue() untuk mengambil data
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validasi inputan
	if name == "" || username == "" || password == "" {
		library.BadResponse(w, "Nama, Username, dan Password tidak boleh kosong")
		return
	}

	// Buat user
	user := model.User{Name: name, Username: username, Password: password}

	var users map[string]interface{}
	users, err := uh.serviceUsers.InputDataUser(user.Name, user.Username, user.Password)
	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	templates.ExecuteTemplate(w, "register-view", users)

}
