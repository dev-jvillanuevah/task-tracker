package taskmanager

import "time"

type Client struct {
	Tasks []Task
}

func NewClient() *Client {
	// TODO: read from the json file here
	return &Client{make([]Task, 0)}
}

func (t *Client) Add(description string) int {
	id := len(t.Tasks) + 1 // starts at 1
	t.Tasks = append(t.Tasks, Task{
		ID:          id,
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now(),
	})
	return id
}

func (t *Client) GetTasks() []Task {
	return t.Tasks
}
