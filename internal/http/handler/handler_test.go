package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"service/internal/models"
	"service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockTaskStorage := new(service.MockTaskStorage)
	taskService := service.NewTaskService(mockTaskStorage)
	handler := NewHandler(taskService)

	task := &models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	mockTaskStorage.On("CreateTask", task).Return("12345", nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/tasks", handler.CreateTask)
	taskJSON, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "12345")
	mockTaskStorage.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockTaskStorage := new(service.MockTaskStorage)
	taskService := service.NewTaskService(mockTaskStorage)
	handler := NewHandler(taskService)

	taskID := "12345"
	expectedTask := &models.Task{
		Id:          taskID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}
	mockTaskStorage.On("GetTask", taskID).Return(expectedTask, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/tasks/:id", handler.GetTask)

	req := httptest.NewRequest(http.MethodGet, "/tasks/12345", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), taskID)
	mockTaskStorage.AssertExpectations(t)
}

func TestListOfTask(t *testing.T) {
	mockTaskStorage := new(service.MockTaskStorage)
	taskService := service.NewTaskService(mockTaskStorage)
	handler := NewHandler(taskService)

	tasks := []*models.Task{
		{Id: "1", Title: "Test Task 1", Description: "Test Description 1", Status: "pending"},
		{Id: "2", Title: "Test Task 2", Description: "Test Description 2", Status: "completed"},
	}
	mockTaskStorage.On("ListOfTask").Return(tasks, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/tasks", handler.ListOfTask)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var result []*models.Task
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Len(t, result, len(tasks))

	for i, task := range tasks {
		assert.Equal(t, task.Id, result[i].Id)
		assert.Equal(t, task.Title, result[i].Title)
		assert.Equal(t, task.Description, result[i].Description)
		assert.Equal(t, task.Status, result[i].Status)
	}

	mockTaskStorage.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockTaskStorage := new(service.MockTaskStorage)
	taskService := service.NewTaskService(mockTaskStorage)
	handler := NewHandler(taskService)

	taskID := "12345"
	task := &models.Task{
		Id:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
	}
	mockTaskStorage.On("UpdateTask", task).Return("Task updated successfully", nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/tasks/:id", handler.UpdateTask)

	taskJSON := `{"title":"Updated Task","description":"Updated Description","status":"completed"}`
	req := httptest.NewRequest(http.MethodPut, "/tasks/12345", strings.NewReader(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task updated successfully")
	mockTaskStorage.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockTaskStorage := new(service.MockTaskStorage)
	taskService := service.NewTaskService(mockTaskStorage)
	handler := NewHandler(taskService)

	taskID := "12345"
	mockTaskStorage.On("DeleteTask", taskID).Return("Task deleted successfully", nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/tasks/:id", handler.DeleteTask)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/12345", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task deleted successfully")
	mockTaskStorage.AssertExpectations(t)
}
