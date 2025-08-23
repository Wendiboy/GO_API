package userService

import (
	"GO_API/internal/taskService"
)

type User struct {
	Id       string             `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks"`
}

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
