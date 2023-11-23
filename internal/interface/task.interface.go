package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type TaskControllerInterface interface {
	Send(ginContext *gin.Context)      // Send a task to the server
	Receive(ginContext *gin.Context)   // Receive a task from the server
	ListTasks(ginContext *gin.Context) // List tasks
}

type TaskServiceInterface interface {
	Send(task model.TaskDTO) error       // Send a task to the server
	Receive(task model.TaskDTO) error    // Receive a task from the server
	ListTasks() ([]model.TaskDTO, error) // List tasks
}
