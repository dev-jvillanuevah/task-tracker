package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dev-jvillanuevah/task-tracker/common"
	"github.com/dev-jvillanuevah/task-tracker/filemanager"
	"github.com/dev-jvillanuevah/task-tracker/taskmanager"
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
		fmt.Print("task-cli ")
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
		case common.CommandDelete:
			err = taskManager.DeleteTask(userCommand.GetID())
		case common.CommandList:
			taskManager.ListTasks()
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("Application terminated")
}
