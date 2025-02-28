package taskmanager

import (
	"fmt"
	"strconv"
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
	case common.CommandUpdate:
		description := strings.Join(s[2:], " ")
		description = strings.ReplaceAll(description, `"`, "")
		id, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, fmt.Errorf("invalid id %s", s[1])
		}
		userCommand.Input1 = id
		userCommand.Description = &description
	case common.CommandList:
	case common.CommandExit:
	default:
		return nil, fmt.Errorf("invalid command: %s", command)
	}
	return userCommand, nil
}
