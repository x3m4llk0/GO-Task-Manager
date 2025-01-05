package main

import (
	"github.com/x3m4llk0/GO-Task-Manager/internal/api"
	"github.com/x3m4llk0/GO-Task-Manager/internal/db"
	"github.com/x3m4llk0/GO-Task-Manager/internal/manager"
	"log"
)

func main() {
	// Инициализация PostgreSQL
	database := db.InitPostgres()
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("Error closing the database: %v", err)
		}
	}()

	// Создаём менеджер задач
	taskManager := manager.NewTaskManager()

	// Настраиваем роутер
	router := api.SetupRouter(taskManager)

	log.Println("Starting server on 0.0.0.0:8080")
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
