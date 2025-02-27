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
		description = strings.ReplaceAll(description, `"`, "")
		userCommand.Input1 = description
	case CommandList:
	case CommandExit:
	default:
		return nil, errors.New(fmt.Sprintf("invalid command: %s", command))
	}
	return userCommand, nil
}
