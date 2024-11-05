package middleware

import (
	"day-22/model"
	"encoding/json"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header
		authHeader := r.Header.Get("Token")

		// Cek jika token kosong atau nilai token adalah "nil"
		if authHeader == "" || authHeader == "nil" {
			// Kirim respons Unauthorized jika token tidak valid
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			w.WriteHeader(http.StatusUnauthorized) // Set status code
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		// Lanjutkan ke handler jika token valid
		next.ServeHTTP(w, r)
	})
}
