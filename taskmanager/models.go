package taskmanager

import "time"

type TaskStatus string

const (
	StatusToDo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (task *Task) ValidStatus() bool {
	return task.Status == StatusToDo || task.Status == StatusInProgress || task.Status == StatusDone
}
