package main

import (
	"day-22/database"
	"day-22/handler"
	"day-22/repository"
	"day-22/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Inisialisasi koneksi ke database
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inisialisasi repository, service, dan handler untuk User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(*userService)

	// Inisialisasi repository, service, dan handler untuk Task
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(*taskService)

	r := chi.NewRouter()

	// Rute untuk User
	r.Get("/users", userHandler.GetAllUsersHandler)
	r.Get("/users/{id}", userHandler.GetUserDetailsHandler)
	r.Post("/create", userHandler.CreateUserHandler)
	r.Post("/login", userHandler.UserLoginHandler)

	// Rute untuk Task

	r.Get("/users/list-task", taskHandler.GetAllTasksHandler)
	r.Post("/users/create-task", taskHandler.CreateTaskHandler)
	r.Get("/users/list-task/{id}", taskHandler.UpdateTaskStatusHandler)

	r.Get("/", handler.FormRegist)

	// Start server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
