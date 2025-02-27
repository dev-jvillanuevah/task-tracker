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

type UserCommand struct {
	Command     TrackCommand
	Input1      any // can be string, int or task status
	Description *string
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

func (userCommand *UserCommand) GetDescription() string {
	if userCommand.Command == CommandUpdate {
		return *userCommand.Description
	}
	return userCommand.Input1.(string)
}
