package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"task-manager/controllers"
	"task-manager/services"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_CreateTask(t *testing.T) {
	taskService := &services.TaskService{}
	taskController := &controllers.TaskController{TaskService: taskService}

	t.Run("ValidRequest", func(t *testing.T) {
		// Create a request body with valid title and description
		requestBody := `{"title": "Task 1", "description": "Description 1"}`

		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodPost, "/api/tasks", strings.NewReader(requestBody))

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the CreateTask handler function
		taskController.CreateTask(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusCreated, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Task created successfully", response["message"])
		assert.NotNil(t, response["data"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a request body with missing title field
		requestBody := `{"description": "Description 2"}`

		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodPost, "/api/tasks", strings.NewReader(requestBody))

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the CreateTask handler function
		taskController.CreateTask(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "title and description are required", response["message"])
		assert.Nil(t, response["data"])
	})
}

func TestTaskController_GetTasks(t *testing.T) {
	taskService := &services.TaskService{}
	taskController := &controllers.TaskController{TaskService: taskService}

	t.Run("ValidRequest", func(t *testing.T) {
		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodGet, "/api/tasks?page=1&limit=10&status=completed&title=Task", nil)

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the GetTasks handler function
		taskController.GetTasks(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Tasks retrieved successfully", response["message"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a new HTTP request with invalid page and limit values
		req, _ := http.NewRequest(http.MethodGet, "/api/tasks?page=abc&limit=xyz&status=completed&title=Task", nil)

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the GetTasks handler function
		taskController.GetTasks(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Tasks retrieved successfully", response["message"])
	})
}

func TestTaskController_GetTaskByID(t *testing.T) {
	taskService := &services.TaskService{}
	taskController := &controllers.TaskController{TaskService: taskService}

	t.Run("ValidRequest", func(t *testing.T) {

		// Create a new HTTP request with a valid task ID
		req, _ := http.NewRequest(http.MethodGet, "/api/tasks/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the GetTaskByID handler function
		taskController.GetTaskByID(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Task retrieved successfully", response["message"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a new HTTP request with an invalid task ID
		req, _ := http.NewRequest(http.MethodGet, "/api/tasks/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the GetTaskByID handler function
		taskController.GetTaskByID(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid task ID", response["message"])
	})
}

func TestTaskController_UpdateTask(t *testing.T) {
	taskService := &services.TaskService{}
	taskController := &controllers.TaskController{TaskService: taskService}

	t.Run("ValidRequest", func(t *testing.T) {
		// Create a request body with valid title, description, and status
		requestBody := `{"title": "Task 1 Updated", "description": "Description 1 Updated", "status": "IN_PROGRESS"}`

		// Create a new HTTP request with a valid task ID
		req, _ := http.NewRequest(http.MethodPut, "/api/tasks/1", strings.NewReader(requestBody))
		req = mux.SetURLVars(req, map[string]string{"id": "1"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the UpdateTask handler function
		taskController.UpdateTask(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Task updated successfully", response["message"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a request body with missing title field
		requestBody := `{"description": "Description 2 Updated", "status": "in_progress"}`

		// Create a new HTTP request with an invalid task ID
		req, _ := http.NewRequest(http.MethodPut, "/api/tasks/abc", strings.NewReader(requestBody))
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the UpdateTask handler function
		taskController.UpdateTask(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid task ID", response["message"])

	})
}

func TestTaskController_DeleteTask(t *testing.T) {
	taskService := &services.TaskService{}
	taskController := &controllers.TaskController{TaskService: taskService}

	t.Run("ValidRequest", func(t *testing.T) {
		// Create a new HTTP request with a valid task ID
		req, _ := http.NewRequest(http.MethodDelete, "/api/tasks/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the DeleteTask handler function
		taskController.DeleteTask(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Task deleted successfully", response["message"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a new HTTP request with an invalid task ID
		req, _ := http.NewRequest(http.MethodDelete, "/api/tasks/abc", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the DeleteTask handler function
		taskController.DeleteTask(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid task ID", response["message"])
		assert.Nil(t, response["data"])
	})
}

func TestTaskController_MarkTaskAsComplete(t *testing.T) {
	taskService := &services.TaskService{}
	taskController := &controllers.TaskController{TaskService: taskService}

	t.Run("ValidRequest", func(t *testing.T) {
		// Create a new HTTP request with a valid task ID
		req, _ := http.NewRequest(http.MethodPut, "/api/tasks/1/complete", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the MarkTaskAsComplete handler function
		taskController.MarkTaskAsComplete(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Task marked as complete", response["message"])
		assert.Nil(t, response["data"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a new HTTP request with an invalid task ID
		req, _ := http.NewRequest(http.MethodPut, "/api/tasks/abc/complete", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "abc"})

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the MarkTaskAsComplete handler function
		taskController.MarkTaskAsComplete(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid task ID", response["message"])
		assert.Nil(t, response["data"])
	})
}
