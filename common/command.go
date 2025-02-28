package common

type UserCommand struct {
	Command     TrackCommand
	Input1      any // can be string, int or task status
	Description *string
}

func (u *UserCommand) GetDescription() string {
	if u.Command == CommandUpdate {
		return *u.Description
	}
	return u.Input1.(string)
}

func (u *UserCommand) GetID() int {
	return u.Input1.(int)
}

func (u *UserCommand) GetTaskStatus() *TaskStatus {
	if u.Input1 == nil {
		return nil
	}
	s := u.Input1.(TaskStatus)
	return &s
}
