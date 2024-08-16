package http

import (
	"service/internal/http/handler"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

func NewGin(service *service.Service) *gin.Engine {

	r := gin.Default()

	handler := handler.NewHandler(service)

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks/:id", handler.GetTask)
	r.GET("/tasks", handler.ListOfTask)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	return r
}