package userService

import "github.com/google/uuid"

type UserService interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserById(id string) (User, error)
	UpdateUser(updatedUser User) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user User) (User, error) {
	user.Id = uuid.NewString()
	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById(id string) (User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return User{}, err
	}

	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(updatedUser User) (User, error) {
	user, err := s.repo.GetUserById(updatedUser.Id)
	if err != nil {
		return User{}, err
	}

	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(id)
}
