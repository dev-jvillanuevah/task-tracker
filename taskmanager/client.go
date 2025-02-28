package taskmanager

import (
	"fmt"
	"time"

	"github.com/dev-jvillanuevah/task-tracker/common"
	"github.com/dev-jvillanuevah/task-tracker/filemanager"
)

type Client struct {
	tasks       []*common.Task
	fileManager *filemanager.Client
}

func NewClient(fileManager *filemanager.Client) *Client {
	return &Client{
		tasks:       []*common.Task{},
		fileManager: fileManager,
	}
}

func (t *Client) AddTask(description string) error {
	id := len(t.tasks) + 1 // starts at 1
	t.tasks = append(t.tasks, &common.Task{
		ID:          id,
		Description: description,
		Status:      common.StatusToDo,
		CreatedAt:   time.Now(),
	})
	err := t.saveTasks()
	if err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	fmt.Printf("Task added successfully (%d)\n", id)
	return nil
}

func (t *Client) UpdateTask(id int, description string) error {
	task := t.findTask(id)
	task.Description = description
	err := t.saveTasks()
	if err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	fmt.Printf("Task updated successfully (%d)\n", id)
	return nil
}

func (t *Client) findTask(id int) *common.Task {
	for _, task := range t.tasks {
		if task.ID == id {
			return task
		}
	}
	return nil
}

func (t *Client) DeleteTask(id int) error {
	t.filterTasks(id)
	err := t.saveTasks()
	if err != nil {
		return fmt.Errorf("error saving tasks: %w", err)
	}
	fmt.Printf("Task deleted successfully (%d)\n", id)
	return nil
}

func (t *Client) filterTasks(id int) {
	var newTasks []*common.Task
	for _, task := range t.tasks {
		if task.ID == id {
			continue
		}
		newTasks = append(newTasks, task)
	}
	t.tasks = newTasks
}

func (t *Client) GetTasks() []*common.Task {
	return t.tasks
}

func (t *Client) ListTasks() {
	for _, task := range t.tasks {
		fmt.Printf("%d - %s - %s\n", task.ID, task.Description, task.GetStatus())
	}
}

func (t *Client) saveTasks() error {
	err := t.fileManager.WriteFile(t.tasks)
	if err != nil {
		return fmt.Errorf("file manager error: %w\n", err)
	}
	return nil
}

func (t *Client) LoadTasks() error {
	err := t.fileManager.ReadFile(&t.tasks)
	if err != nil {
		return fmt.Errorf("file manager error: %w\n", err)
	}
	return nil
}
