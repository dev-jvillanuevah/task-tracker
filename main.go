package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dev-jvillanuevah/task-tracker/taskmanager"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// start taskmanager client
	taskManagerClient := taskmanager.NewClient()
	for {
		fmt.Print("task-cli ")
		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())
		if command == "exit" {
			fmt.Println("Exiting application...")
			break
		}
		if command == "add" {
			taskManagerClient.Add("some description")
		}
		fmt.Printf("tasks: %+v\n", taskManagerClient.GetTasks())
	}
	fmt.Println("Application terminated")
}
