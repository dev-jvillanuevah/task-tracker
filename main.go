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
	err := taskManager.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Print("task-cli ")
		scanner.Scan()
		var userCommand *common.UserCommand
		userCommand, err = taskmanager.ParseUserInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			break
		}
		if userCommand.Command == common.CommandExit {
			fmt.Println("Exiting application...")
			break
		}
		switch userCommand.Command {
		case common.CommandAdd:
			_, addErr := taskManager.AddTask(userCommand.GetDescription())
			if addErr != nil {
				fmt.Println(addErr)
				break
			}
		case common.CommandUpdate:
			updateErr := taskManager.UpdateTask(
				userCommand.GetID(),
				userCommand.GetDescription(),
			)
			if updateErr != nil {
				fmt.Println(updateErr)
				break
			}
		case common.CommandList:
			taskManager.ListTasks()
		}
	}
	fmt.Println("Application terminated")
}
