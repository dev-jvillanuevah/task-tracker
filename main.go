package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dev-jvillanuevah/task-tracker/taskmanager"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	client := taskmanager.NewClient()
	for {
		fmt.Print("task-cli ")
		scanner.Scan()
		userCommand, err := taskmanager.ParseUserInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			break
		}
		if userCommand.Command == taskmanager.CommandExit {
			fmt.Println("Exiting application...")
			break
		}
		if userCommand.Command == taskmanager.CommandAdd {
			client.Add(userCommand.GetDescription())
		}
		err = client.WriteFile()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("Application terminated")
}
