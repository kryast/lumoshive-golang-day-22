package middleware

import (
	"day-22/model"
	"day-22/repository"
	"encoding/json"
	"net/http"
)

type Middleware struct {
	RepoUser repository.UserRepositoryDB
}

func NewMiddleware(repo repository.UserRepositoryDB) *Middleware {
	return &Middleware{RepoUser: repo}
}

func (mw *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		authHeader := r.Header.Get("Token")

		// Jika token kosong atau "nil", langsung return Unauthorized
		if authHeader == "" || authHeader == "nil" {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized: Token tidak ditemukan",
				Data:       nil,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Verifikasi token di database
		user, err := mw.RepoUser.GetUserByToken(authHeader)
		if err != nil || user.Token != authHeader {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized: Token tidak valid",
				Data:       nil,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
