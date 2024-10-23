package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InMemoryTaskManager struct {
	tasks  []Task
	nextID int
}

func NewInMemoryTaskManager() *InMemoryTaskManager {
	return &InMemoryTaskManager{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (m *InMemoryTaskManager) AddTask(c *gin.Context) {
	var newTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task := Task{
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

func (m *InMemoryTaskManager) ListTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, m.tasks)
}

func (m *InMemoryTaskManager) UpdateTaskStatus(c *gin.Context) {
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

func (m *InMemoryTaskManager) GetTaskById(c *gin.Context) {
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

func (m *InMemoryTaskManager) RemoveTask(c *gin.Context) {
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

func main() {
	manager := NewInMemoryTaskManager()
	router := gin.Default()

	router.POST("/task", manager.AddTask)
	router.GET("/tasks", manager.ListTasks)
	router.GET("/task", manager.GetTaskById)
	router.PATCH("/task/:id", manager.UpdateTaskStatus)
	router.DELETE("/task/:id", manager.RemoveTask)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
