package handlers

import (
	"GO_API/internal/taskService"
	"GO_API/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandlers(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.Id,
			Task:   &tsk.TaskBody,
			IsDone: &tsk.Is_done,
		}
		response = append(response, task)
	}

	return response, nil

}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		TaskBody: *taskRequest.Task,
		Is_done:  *taskRequest.IsDone,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.Id,
		Task:   &createdTask.TaskBody,
		IsDone: &createdTask.Is_done,
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	taskRequest := request.Body

	taskToUpdate := taskService.Task{
		Id:       id,
		TaskBody: *taskRequest.Task,
		Is_done:  *taskRequest.IsDone,
	}

	updatedTask, err := h.service.UpdateTask(taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId202JSONResponse{
		Task:   &updatedTask.TaskBody,
		IsDone: &updatedTask.Is_done,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := h.service.DeleteTask(id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
