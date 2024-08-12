package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-manager/models"
	"task-manager/services"
	"task-manager/utils"

	"github.com/gorilla/mux"
)

// TaskController handles task-related HTTP requests.
type TaskController struct {
	TaskService *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{TaskService: service}
}

// CreateTask creates a new task.
// It expects a JSON payload with "title" and "description" fields.
// On success, it returns the created task in the response.
func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid request", nil)
		return
	}

	if input.Title == "" || input.Description == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "title and description are required", nil)
		return
	}

	task := tc.TaskService.CreateTask(input.Title, input.Description)
	utils.SendJSONResponse(w, http.StatusCreated, "success", "Task created successfully", task)
}

// GetTasks retrieves tasks based on query parameters.
// It supports "page", "limit", "status", and "title" query parameters.
// On success, it returns the list of tasks in the response.
func (tc *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	status := models.Status(r.URL.Query().Get("status"))
	title := r.URL.Query().Get("title")

	page := 1
	limit := 10

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	tasks := tc.TaskService.GetTasks(page, limit, status, title)
	utils.SendJSONResponse(w, http.StatusOK, "success", "Tasks retrieved successfully", tasks)
}

// GetTaskByID retrieves a task by ID.
// It expects the task ID as a URL parameter.
// On success, it returns the task in the response.
func (tc *TaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid task ID", nil)
		return
	}
	task, err := tc.TaskService.GetTaskByID(id)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, "error", "Task not found", nil)
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, "success", "Task retrieved successfully", task)
}

// UpdateTask updates a task by ID.
// It expects the task ID as a URL parameter and a JSON payload with "title", "description", and "status" fields.
// On success, it returns a success message in the response.
func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid task ID", nil)
		return
	}
	var input struct {
		Title       string        `json:"title"`
		Description string        `json:"description"`
		Status      models.Status `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid request", nil)
		return
	}
	if err := tc.TaskService.UpdateTask(id, input.Title, input.Description, input.Status); err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, "error", "Task not found", nil)
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, "success", "Task updated successfully", nil)
}

// DeleteTask deletes a task by ID.
// It expects the task ID as a URL parameter.
// On success, it returns a success message in the response.
func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid task ID", nil)
		return
	}
	if err := tc.TaskService.DeleteTask(id); err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, "error", "Task not found", nil)
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, "success", "Task deleted successfully", nil)
}

// MarkTaskAsComplete marks a task as complete.
// It expects the task ID as a URL parameter.
// On success, it returns a success message in the response.
func (tc *TaskController) MarkTaskAsComplete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid task ID", nil)
		return
	}
	if err := tc.TaskService.MarkTaskAsComplete(id); err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, "error", "Task not found", nil)
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, "success", "Task marked as complete", nil)
}
