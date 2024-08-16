package service

import (
	"service/internal/repository"
	models "service/internal/models"
)

type Service struct {
	TaskService repository.ITaskStorage
}

func NewTaskService(service repository.ITaskStorage) *Service  {
	return &Service{
		TaskService: service,
	}
}

func (s *Service) CreateTask(req *models.Task) (string, error) {
	return s.TaskService.CreateTask(req)
}

func (s *Service) GetTask(id string) (*models.Task, error) {
	return s.TaskService.GetTask(id)
}

func (s *Service) ListOfTask() ([]*models.Task, error) {
	return s.TaskService.ListOfTask()
}

func (s *Service) UpdateTask(req *models.Task) (string, error) {
	return s.TaskService.UpdateTask(req)
}

func (s *Service) DeleteTask(id string) (string, error) {
	return s.TaskService.DeleteTask(id)
}