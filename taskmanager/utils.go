package taskmanager

import (
	"errors"
	"fmt"
	"strings"
)

func ParseUserInput(input string) (*UserCommand, error) {
	s := strings.Split(input, " ")
	command := TrackCommand(s[0])
	userCommand := &UserCommand{
		Command: command,
	}
	switch command {
	case CommandAdd:
		description := strings.Join(s[1:], " ")
		userCommand.Input1 = description
	case CommandExit:
	default:
		return nil, errors.New(fmt.Sprintf("invalid command: %s\n", command))
	}
	return userCommand, nil
}
