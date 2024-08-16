package repository

import (
	models "service/internal/models"
)

type ITaskStorage interface {
	CreateTask(req *models.Task) (string, error)
	GetTask(id string) (*models.Task, error)
	ListOfTask() ([]*models.Task, error)
	UpdateTask(req *models.Task) (string, error)
	DeleteTask(id string) (string, error)
}
