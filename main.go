package main

import (
	"day-22/handler"
	"fmt"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("POST /create", handler.CreateUserHandler)
	serverMux.HandleFunc("GET /users", handler.GetAllUsersHandler)

	serverMux.HandleFunc("POST /users/create-task", handler.CreateTaskHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", serverMux)

}
