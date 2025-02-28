package taskmanager

import (
	"fmt"
	"time"

	"github.com/dev-jvillanuevah/task-tracker/common"
	"github.com/dev-jvillanuevah/task-tracker/filemanager"
)

type Client struct {
	Tasks       []*common.Task
	FileManager *filemanager.Client
}

func NewClient(fileManager *filemanager.Client) *Client {
	return &Client{
		Tasks:       []*common.Task{},
		FileManager: fileManager,
	}
}

func (t *Client) Add(description string) (int, error) {
	id := len(t.Tasks) + 1 // starts at 1
	t.Tasks = append(t.Tasks, &common.Task{
		ID:          id,
		Description: description,
		Status:      common.StatusToDo,
		CreatedAt:   time.Now(),
	})
	err := t.saveTasks()
	if err != nil {
		return 0, fmt.Errorf("error saving tasks: %w", err)
	}
	return id, nil
}

func (t *Client) GetTasks() []*common.Task {
	return t.Tasks
}

func (t *Client) ListTasks() {
	for _, task := range t.Tasks {
		fmt.Printf("%s - %s\n", task.Description, task.GetStatus())
	}
}

func (t *Client) saveTasks() error {
	err := t.FileManager.WriteFile(t.Tasks)
	if err != nil {
		return fmt.Errorf("file manager error: %w\n", err)
	}
	return nil
}

func (t *Client) LoadTasks() error {
	err := t.FileManager.ReadFile(&t.Tasks)
	if err != nil {
		return fmt.Errorf("file manager error: %w\n", err)
	}
	return nil
}
