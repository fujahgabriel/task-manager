package main

import (
	"net/http"
	"task-manager/controllers"
	"task-manager/routes"
	"task-manager/services"

	"github.com/gorilla/mux"
)

func main() {
	taskService := services.NewTaskService()
	userService := services.NewUserService()
	taskController := &controllers.TaskController{TaskService: taskService}
	userController := &controllers.UserController{UserService: userService}

	router := mux.NewRouter()

	// User authentication routes
	router.HandleFunc("/api/register", userController.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", userController.Login).Methods(http.MethodPost)

	// Task management routes
	routes.RegisterTaskRoutes(router, taskController)

	http.ListenAndServe(":8080", router)
}
