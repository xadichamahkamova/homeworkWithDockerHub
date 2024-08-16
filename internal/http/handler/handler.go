package handler

import (
	"service/internal/service"

	"github.com/gin-gonic/gin"
	models "service/internal/models"
)

type HandlerST struct {
	Service service.Service
}

func NewHandler(service service.Service) *HandlerST {
	return &HandlerST{
		Service: service,
	}
}

func (h *HandlerST) CreateTask(c *gin.Context) {

	req := models.Task{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message":err.Error()})
		return
	}
	resp, err := h.Service.CreateTask(&req)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	c.JSON(200, gin.H{"id":resp})
}

func (h *HandlerST) GetTask(c *gin.Context) {

	id := c.Param("id")
	resp, err := h.Service.GetTask(id)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) ListOfTask(c *gin.Context) {

	resp, err := h.Service.ListOfTask()
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) UpdateTask(c *gin.Context) {

	req := models.Task{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message":err.Error()})
		return
	}

	req.Id = c.Param("id")
	resp, err := h.Service.UpdateTask(&req)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	c.JSON(200, gin.H{"message":resp})
}

func (h *HandlerST) DeleteTask(c *gin.Context) {

	id := c.Param("id")
	resp, err := h.Service.DeleteTask(id)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	c.JSON(200, gin.H{"message":resp})
}