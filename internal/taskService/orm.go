package taskService

type Task struct {
	ID       string
	TaskBody string
	Is_done  bool
}

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
