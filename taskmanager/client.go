package taskmanager

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Client struct {
	Tasks []*Task
}

func NewClient() *Client {
	// TODO: read from the json file here
	return &Client{Tasks: []*Task{}}
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

func (t *Client) WriteFile() error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return fmt.Errorf("error creating tasks.json: %w", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(t.Tasks); err != nil {
		return fmt.Errorf("error writing tasks.json: %w", err)
	}
	return nil
}
