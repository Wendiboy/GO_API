package handlers

import (
	"GO_API/internal/userService"
	"GO_API/internal/web/users"
	"context"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandlers(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()

	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, v := range allUsers {
		user := users.User{
			Id:       &v.Id,
			Email:    &v.Email,
			Password: &v.Password,
		}
		response = append(response, user)

	}

	return response, nil
}

func (u *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := u.service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.Id,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil

}

func (u *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body

	userToUpdate := userService.User{
		Id:       id,
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := u.service.UpdateUser(userToUpdate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &id,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

func (u *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	err := u.service.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	response := users.DeleteUsersId204Response{}

	return response, nil
}
