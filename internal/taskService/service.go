package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(updatedTask Task) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s taskService) CreateTask(task Task) (Task, error) {
	task.Id = uuid.NewString()

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s taskService) GetAllTasks() ([]Task, error) {

	return s.repo.GetAllTasks()
}

func (s taskService) GetTaskById(id string) (Task, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return Task{}, err
	}

	return s.repo.GetTaskById(id)
}

func (s taskService) UpdateTask(updatedTask Task) (Task, error) {

	task, err := s.repo.GetTaskById(updatedTask.Id)

	if err != nil {
		return Task{}, err
	}

	task.TaskBody = updatedTask.TaskBody
	task.Is_done = updatedTask.Is_done

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s taskService) DeleteTask(id string) error {

	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteTask(id)
}
