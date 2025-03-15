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
	CommandHelp           TrackCommand = "help"
)

type TaskStatus string

const (
	StatusToDo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

var ValidStatus = []TaskStatus{StatusToDo, StatusInProgress, StatusDone}

var HelpMessage = `[COMMAND] [ARG] [OPTIONAL ARG]
list of commands
add	[TASK DESCRIPTION] -- adds a new task 
update [TASK ID] [TASK DESCRIPTION] -- updates a task
delete [TASK ID] -- deletes a task
mark-in-progress [TASK ID] -- updates the status of a task to in-progress
mark-done [TASK ID]	-- updates the status of a task to done
list -- list all tasks
list [TASK STATUS] -- list tasks with the selected status -- [todo|in-progress|done]`
