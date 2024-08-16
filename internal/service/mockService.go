package service

import (
	"service/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockTaskStorage struct {
	mock.Mock
}

func (m *MockTaskStorage) CreateTask(req *models.Task) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func (m *MockTaskStorage) GetTask(id string) (*models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskStorage) ListOfTask() ([]*models.Task, error) {
	args := m.Called()
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockTaskStorage) UpdateTask(req *models.Task) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func (m *MockTaskStorage) DeleteTask(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}