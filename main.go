package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	middleware "github.com/labstack/echo/v4/middleware"
)

var db *gorm.DB //переменная для работы с БД

func initDB() {
	// Функция инициализации БД
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable" //Data Source Name
	var err error

	// Подключаемся к БД, если не удалось выдаем fatal
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v ", err)
	}

	// Запускаем автомиграцию на основе структуры Task
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v ", err)
	}
}

type Task struct {
	ID       string
	TaskBody string
	Is_done  bool
}

type requestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func postTask(c echo.Context) error {
	req := new(requestBody)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	task := Task{
		ID:       uuid.NewString(),
		TaskBody: req.Task,
		Is_done:  req.IsDone,
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add task"})
	}

	return c.JSON(http.StatusOK, task)
}

func getTasks(c echo.Context) error {
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid request"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func patchTask(c echo.Context) error {
	var task Task

	id := c.Param("id")

	req := new(requestBody)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.Find(&task, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not find task"})
	}

	task.TaskBody = req.Task
	task.Is_done = req.IsDone

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	var task Task

	id := c.Param("id")

	if err := db.Delete(&task, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/tasks", postTask)
	e.GET("/tasks", getTasks)
	e.PATCH("/tasks/:id", patchTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Start("localhost:8080")
}
