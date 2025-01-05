package manager

import (
	"github.com/x3m4llk0/GO-Task-Manager/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// TaskManager управляет задачами в памяти.
type TaskManager struct {
	tasks  []models.Task
	nextID int
}

// NewTaskManager создаёт новый менеджер задач.
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

// AddTask добавляет новую задачу.
func (m *TaskManager) AddTask(c *gin.Context) {
	var newTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task := models.Task{
		ID:          m.nextID,
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.tasks = append(m.tasks, task)
	m.nextID++
	c.IndentedJSON(http.StatusCreated, task)
}

// ListTasks возвращает все задачи.
func (m *TaskManager) ListTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, m.tasks)
}

// UpdateTaskStatus обновляет статус задачи.
func (m *TaskManager) UpdateTaskStatus(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID should be an integer"})
		return
	}

	var updateData struct {
		Completed bool `json:"completed"`
	}

	if err := c.BindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	for i, task := range m.tasks {
		if task.ID == intID {
			m.tasks[i].Completed = updateData.Completed
			m.tasks[i].UpdatedAt = time.Now()
			c.IndentedJSON(http.StatusOK, m.tasks[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// GetTaskByID возвращает задачу по ID.
func (m *TaskManager) GetTaskByID(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID query parameter is missing"})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID should be an integer"})
		return
	}

	for _, task := range m.tasks {
		if task.ID == intID {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// RemoveTask удаляет задачу по ID.
func (m *TaskManager) RemoveTask(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID should be an integer"})
		return
	}

	for i, task := range m.tasks {
		if task.ID == intID {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
