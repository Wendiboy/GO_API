package handlers

import (
	"GO_API/internal/taskService"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandlers(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) PostTask(c echo.Context) error {

	var req taskService.RequestBody

	fmt.Println(c)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	task, err := h.service.CreateTask(req.Task, req.IsDone)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PatchTask(c echo.Context) error {

	id := c.Param("id")

	var req taskService.RequestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updatedTask, err := h.service.UpdateTask(id, req.Task, req.IsDone)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {

	id := c.Param("id")

	err := h.service.DeleteTask(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
