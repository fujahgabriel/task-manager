package services

import (
	"errors"
	"strings"
	"sync"
	"task-manager/models"
)

type TaskService struct {
	tasks  []models.Task
	mutex  sync.Mutex
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (s *TaskService) CreateTask(title, description string) models.Task {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task := models.Task{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		Status:      models.Todo,
	}
	s.tasks = append(s.tasks, task)
	s.nextID++
	return task
}

func (s *TaskService) GetTasks(page, pageSize int, status models.Status, title string) []models.Task {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var filteredTasks []models.Task

	// Filter tasks based on status and title
	for _, task := range s.tasks {
		if (status == "" || task.Status == status) && (title == "" || strings.Contains(strings.ToLower(task.Title), strings.ToLower(title))) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	// Implement pagination
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(filteredTasks) {
		return []models.Task{}
	}

	if end > len(filteredTasks) {
		end = len(filteredTasks)
	}

	return filteredTasks[start:end]
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) findAndUpdateTask(id int, updateFunc func(*models.Task)) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			updateFunc(&s.tasks[i])
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) UpdateTask(id int, title string, description string, status models.Status) error {
	return s.findAndUpdateTask(id, func(task *models.Task) {
		task.Title = title
		task.Description = description
		task.Status = status
	})
}

func (s *TaskService) MarkTaskAsComplete(id int) error {
	return s.findAndUpdateTask(id, func(task *models.Task) {
		task.Status = models.Completed
	})
}

func (s *TaskService) DeleteTask(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
