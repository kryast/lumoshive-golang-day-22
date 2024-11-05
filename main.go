package main

import (
	"day-22/database"
	"day-22/handler"
	"day-22/middleware"
	"day-22/repository"
	"day-22/service"
	"fmt"
	"net/http"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewCustomerHandler(*userService)

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(*taskService)

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("POST /create", userHandler.CreateUserHandler)
	serverMux.HandleFunc("POST /login", userHandler.UserLoginHandler)

	serverMux.Handle("GET /users", middleware.Middleware(http.HandlerFunc(userHandler.GetAllUsersHandler)))

	serverMux.Handle("POST /users/create-task", middleware.Middleware(http.HandlerFunc(taskHandler.CreateTaskHandler)))

	serverMux.Handle("GET /users/list-task", middleware.Middleware(http.HandlerFunc(taskHandler.GetAllTasksHandler)))

	serverMux.Handle("POST /users/list-task/{id}", middleware.Middleware(http.HandlerFunc(taskHandler.UpdateTaskStatusHandler)))

	serverMux.Handle("GET /users/{id}", middleware.Middleware(http.HandlerFunc(userHandler.GetUserDetailsHandler)))

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", serverMux)

}
