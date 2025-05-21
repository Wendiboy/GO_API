package main

import (
	"GO_API/internal/db"
	"GO_API/internal/handlers"

	"GO_API/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"

	middleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DataBase: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandlers(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/tasks", taskHandlers.PostTask)
	e.GET("/tasks", taskHandlers.GetTasks)
	e.PATCH("/tasks/:id", taskHandlers.PatchTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
