package taskmanager

import (
	"fmt"
	"time"

	"github.com/dev-jvillanuevah/task-tracker/filemanager"
)

type Client struct {
	Tasks       []*Task
	FileManager *filemanager.Client
}

func NewClient(fileManager *filemanager.Client) *Client {
	return &Client{
		Tasks:       []*Task{},
		FileManager: fileManager,
	}
}

func (t *Client) Add(description string) int {
	id := len(t.Tasks) + 1 // starts at 1
	t.Tasks = append(t.Tasks, &Task{
		ID:          id,
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now(),
	})
	return id
}

func (t *Client) GetTasks() []*Task {
	return t.Tasks
}

func (t *Client) ListTasks() {
	for _, task := range t.Tasks {
		fmt.Printf("%s - %s\n", task.Description, task.GetStatus())
	}
}

func (t *Client) SaveTasks() error {
	err := t.FileManager.WriteFile(t.Tasks)
	if err != nil {
		return fmt.Errorf("file manager error: %w\n", err)
	}
	return nil
}

func (t *Client) LoadTasks() error {
	err := t.FileManager.ReadFile(t.Tasks)
	if err != nil {
		return fmt.Errorf("file manager error: %w\n", err)
	}
	return nil
}
