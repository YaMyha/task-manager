package main

import (
	"context"
	"fmt"
	"os"
	"task-manager/internal/storage"
	"task-manager/internal/task"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task add \"task title\"")
		return
	}

	command := os.Args[1]
	title := os.Args[2]

	db, err := pgxpool.New(context.Background(),
		"postgres://task:task@localhost:5434/task_manager")
	if err != nil {
		fmt.Println("DB connection error:", err)
		return
	}

	defer db.Close()

	repo := storage.NewPostgresStorage(db)
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
