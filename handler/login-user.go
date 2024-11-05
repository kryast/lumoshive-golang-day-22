package handler

import (
	"day-22/library"
	"day-22/model"
	"encoding/json"
	"net/http"
)

func (uh *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	userData, err := uh.serviceUsers.LoginUser(user)
	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	library.SuccessResponse(w, "Login success", userData)

}
