package main

import (
	"GO_API/internal/db"
	"GO_API/internal/handlers"
	"GO_API/internal/userService"
	"GO_API/internal/web/tasks"
	"GO_API/internal/web/users"

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

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandlers(taskService)

	userRepo := userService.NewUserRepo(database)
	UserService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandlers(UserService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	strictHandler2 := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictHandler2)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
