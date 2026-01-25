package main

import (
	"fmt"
	"os"
	"task-manager/internal/storage"
	"task-manager/internal/task"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task add \"task title\"")
		return
	}

	command := os.Args[1]
	title := os.Args[2]

	repo := storage.NewJSONStorage("data/tasks.json")
	service := task.NewService(repo)

	switch command {
	case "add":
		if err := service.Add(title); err != nil {
			fmt.Println("Error adding task:", err)
			return
		}

		fmt.Println("Task added successfully")
	default:
		fmt.Println("Unknown command:", command)
	}
}
