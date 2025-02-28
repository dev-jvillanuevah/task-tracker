package common

import (
	"time"
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

func (task *Task) GetStatus() string {
	switch task.Status {
	case StatusToDo:
		return "To Do"
	case StatusDone:
		return "Done"
	case StatusInProgress:
		return "In Progress"
	default:
		return "Unknown"
	}
}
