package handler

import (
	"day-22/database"
	"day-22/repository"
	"day-22/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL parameter (misalnya /user/{id})
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Menggunakan service untuk mendapatkan detail user
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo) // Gantilah dengan repository yang sesuai
	user, err := userService.GetUserDetails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirimkan response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"StatusCode": 200,
		"Message":    "User details fetched successfully",
		"Data":       user,
	})
}
