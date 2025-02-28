package common

type TrackCommand string

const (
	CommandAdd            TrackCommand = "add"
	CommandList           TrackCommand = "list"
	CommandUpdate         TrackCommand = "update"
	CommandMarkInProgress TrackCommand = "mark-in-progress"
	CommandMarkDone       TrackCommand = "mark-done"
	CommandDelete         TrackCommand = "delete"
	CommandExit           TrackCommand = "exit"
)

type TaskStatus string

const (
	StatusToDo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)
