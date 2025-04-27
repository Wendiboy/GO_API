package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

var task string

type requestBody struct {
	Task string `JSON:"task"`
}

func postHandler(c echo.Context) error {
	req := new(requestBody)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	task = req.Task
	fmt.Print(task)
	return c.JSON(http.StatusOK, "OK")
}

func getHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello, "+task)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.POST("/", postHandler)
	e.GET("/", getHandler)
	e.Start("localhost:8080")
}
