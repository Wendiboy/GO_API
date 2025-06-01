package taskService

type Task struct {
	Id       string `json:"id"`
	TaskBody string `json:"taskBody"`
	Is_done  bool   `json:"is_done"`
}

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
