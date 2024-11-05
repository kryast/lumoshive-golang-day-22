package handler

import (
	"day-22/database"
	"day-22/model"
	"day-22/repository"
	"day-22/service"
	"encoding/json"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Error Server",
		Data:       nil,
	}
	if err != nil {
		badResponse.Message = "Invalid input"
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		badResponse.Message = "Database connection error"
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	// Create user repository and service
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	// Call the service to create a user
	var users map[string]interface{}
	users, err = userService.InputDataUser(user.Name, user.Username, user.Password)
	if err != nil {
		badResponse.Message = err.Error()
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	// Success response
	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "User created successfully",
		Data:       users,
	}
	json.NewEncoder(w).Encode(response)
}
