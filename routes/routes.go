package routes

import (
	"net/http"
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(router *mux.Router, taskController *controllers.TaskController) {
	api := router.PathPrefix("/api").Subrouter()
	api.Handle("/tasks", middleware.JWTAuthMiddleware(http.HandlerFunc(taskController.CreateTask))).Methods(http.MethodPost)
	api.Handle("/tasks", middleware.JWTAuthMiddleware(http.HandlerFunc(taskController.GetTasks))).Methods(http.MethodGet)
	api.Handle("/tasks/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(taskController.GetTaskByID))).Methods(http.MethodGet)
	api.Handle("/tasks/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(taskController.UpdateTask))).Methods(http.MethodPut)
	api.Handle("/tasks/{id:[0-9]+}", middleware.JWTAuthMiddleware(http.HandlerFunc(taskController.DeleteTask))).Methods(http.MethodDelete)
	api.Handle("/tasks/{id:[0-9]+}/complete", middleware.JWTAuthMiddleware(http.HandlerFunc(taskController.MarkTaskAsComplete))).Methods(http.MethodPatch)
}
