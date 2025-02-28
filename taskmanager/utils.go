package taskmanager

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dev-jvillanuevah/task-tracker/common"
)

func ParseUserInput(input string) (*common.UserCommand, error) {
	s := strings.Split(input, " ")
	command := common.TrackCommand(s[0])
	userCommand := &common.UserCommand{
		Command: command,
	}
	switch command {
	case common.CommandAdd:
		description := strings.Join(s[1:], " ")
		description = strings.ReplaceAll(description, `"`, "")
		userCommand.Input1 = description
	case common.CommandList:
	case common.CommandExit:
	default:
		return nil, errors.New(fmt.Sprintf("invalid command: %s", command))
	}
	return userCommand, nil
}
