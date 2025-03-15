package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dev-jvillanuevah/task-tracker/common"
	"github.com/dev-jvillanuevah/task-tracker/filemanager"
	"github.com/dev-jvillanuevah/task-tracker/taskmanager"
	"github.com/fatih/color"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fileManager := filemanager.NewClient("tasks.json")
	taskManager := taskmanager.NewClient(fileManager)
	loadTasksErr := taskManager.LoadTasks()
	if loadTasksErr != nil {
		fmt.Println(loadTasksErr)
	}
	for {
		// For green color for "task-cli" text
		color.Set(color.FgGreen)
		fmt.Print("task-cli ")
		color.Unset()
		color.Set(color.FgYellow)
		scanner.Scan()
		userCommand, parseInputErr := taskmanager.ParseUserInput(scanner.Text())
		if parseInputErr != nil {
			fmt.Println(parseInputErr)
			break
		}
		if userCommand.Command == common.CommandExit {
			fmt.Println("Exiting application...")
			break
		}
		var err error
		switch userCommand.Command {
		case common.CommandAdd:
			err = taskManager.AddTask(userCommand.GetDescription())
		case common.CommandUpdate:
			err = taskManager.UpdateTask(userCommand.GetID(), userCommand.GetDescription())
		case common.CommandMarkInProgress:
			err = taskManager.MarkInProgress(userCommand.GetID())
		case common.CommandMarkDone:
			err = taskManager.MarkDone(userCommand.GetID())
		case common.CommandDelete:
			err = taskManager.DeleteTask(userCommand.GetID())
		case common.CommandList:
			err = taskManager.ListTasks(userCommand.GetTaskStatus())
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("Application terminated")
	defer color.Unset()
}
