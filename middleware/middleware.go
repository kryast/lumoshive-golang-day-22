package middleware

import (
	"day-22/database"
	"day-22/model"
	"day-22/repository"
	"day-22/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")

		db, err := database.InitDB()
		if err != nil {
			fmt.Println("err ", err)
		}
		repo := repository.NewUserRepository(db)
		serviceUser := service.NewUserService(repo)

		token := serviceUser.CheckToken(authHeader)
		if token == "" || token == "nil" {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
