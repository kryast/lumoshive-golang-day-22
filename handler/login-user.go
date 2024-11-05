package handler

import (
	"day-22/database"
	"day-22/model"
	"day-22/repository"
	"day-22/service"
	"encoding/json"
	"net/http"
)

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Error Server",
		Data:       nil,
	}
	if err != nil {
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Error Server",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	userData, err := userService.LoginUser(user)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Login failed",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	Response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       userData,
	}
	json.NewEncoder(w).Encode(Response)

}
