package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

type task struct {
	ID       string
	TaskBody string
}

var Tasks = []task{{"3", "fdfdf"}}

type requestBody struct {
	Task string `json:"task"`
}

func postHandler(c echo.Context) error {
	req := new(requestBody)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	Tasks = append(Tasks, task{"1", req.Task})
	return c.JSON(http.StatusOK, "OK")
}

func getHandler(c echo.Context) error {
	fmt.Print(Tasks)
	return c.JSON(http.StatusOK, Tasks)
}

func patchHandler(c echo.Context) error {

	id := c.Param("id")

	req := new(requestBody)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for i, v := range Tasks {
		if v.ID == id {
			Tasks[i].TaskBody = req.Task
		}
	}

	return c.JSON(http.StatusOK, Tasks)
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	for i, v := range Tasks {
		if v.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.POST("/", postHandler)
	e.GET("/", getHandler)
	e.PATCH("/:id", patchHandler)
	e.DELETE("/:id", deleteHandler)
	e.Start("localhost:8080")
}
