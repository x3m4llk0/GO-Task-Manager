package api

import (
	"github.com/gin-gonic/gin"
	"github.com/x3m4llk0/GO-Task-Manager/internal/manager"
)

// SetupRouter настраивает маршруты приложения.
func SetupRouter(taskManager *manager.TaskManager) *gin.Engine {
	router := gin.Default()

	router.POST("/task", taskManager.AddTask)
	router.GET("/tasks", taskManager.ListTasks)
	router.GET("/task", taskManager.GetTaskByID)
	router.PATCH("/task/:id", taskManager.UpdateTaskStatus)
	router.DELETE("/task/:id", taskManager.RemoveTask)

	return router
}
