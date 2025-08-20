package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task [add|list|complete] [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			return
		}
		completeTask(os.Args[2])
	default:
		fmt.Println("Unknown command")
	}
}

func addTask(description string) {
	fmt.Printf("Adding task: %s\n", description)
}

func listTasks() {
	fmt.Println("Listing all tasks")
}

func completeTask(taskID string) {
	fmt.Printf("Completing task %s\n", taskID)
}
