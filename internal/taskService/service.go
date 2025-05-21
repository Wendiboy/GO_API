package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(taskBody string, isDone bool) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(id string, taskBody string, isDone bool) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s taskService) CreateTask(taskBody string, isDone bool) (Task, error) {
	task := Task{
		ID:       uuid.NewString(),
		TaskBody: taskBody,
		Is_done:  isDone,
	}

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

func (s taskService) UpdateTask(id string, taskBody string, isDone bool) (Task, error) {
	task, err := s.repo.GetTaskById(id)

	if err != nil {
		return Task{}, err
	}

	task.TaskBody = taskBody
	task.Is_done = isDone

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
