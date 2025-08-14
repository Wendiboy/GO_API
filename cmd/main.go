package main

import (
	"GO_API/internal/db"
	"GO_API/internal/handlers"
	"GO_API/internal/web/tasks"

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

	repo := taskService.NewTaskRepository(database)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewTaskHandlers(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
