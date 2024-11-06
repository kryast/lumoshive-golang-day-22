package handler

import (
	"fmt"
	"net/http"
)

func (uh *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := uh.serviceUsers.GetAllDataUser()
	if err != nil {
		fmt.Println("err ", err)
	}

	// fmt.Println("data :", *customers)
	templates.ExecuteTemplate(w, "list-user-view", *users)
}
