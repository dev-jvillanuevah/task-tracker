package taskmanager

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
