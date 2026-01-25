# Task Manager

A lightweight, clean command-line task management application built with Go, featuring a layered architecture that separates concerns and promotes maintainability.

## Overview

Task Manager is a CLI tool that allows you to create and manage tasks efficiently. It demonstrates Go best practices including dependency injection, interface-based design, and the Repository pattern for data persistence.

## Project Structure

```
task-manager/
├── cmd/
│   └── task/
│       └── main.go              # Application entry point
├── internal/
│   ├── task/
│   │   ├── model.go             # Task entity definition
│   │   ├── service.go           # Business logic implementation
│   │   └── repository.go        # Repository interface definition
│   └── storage/
│       └── json.go              # JSON storage implementation
├── data/
│   └── tasks.json               # Persistent task data
└── go.mod                        # Go module definition
```

## Architecture

This project follows a clean, layered architecture pattern with clear separation of concerns:

### 1. **Entry Point** (`cmd/task/main.go`)

The main package serves as the application's entry point:
- Parses command-line arguments
- Initializes dependencies (storage and service)
- Routes commands to appropriate handlers
- Does **not** contain business logic, database operations, or domain logic

This keeps the main function focused on orchestration rather than implementation details.

### 2. **Domain Layer** (`internal/task/`)

The task package contains the core business logic and data models:

#### `model.go` - Task Entity
Defines the `Task` struct representing a task entity with persistence:
```go
type Task struct {
    ID    int    // Unique identifier
    Title string // Task description
    Done  bool   // Completion status
}
```

#### `repository.go` - Data Access Interface
Defines the `Repository` interface that abstracts data persistence operations:
```go
type Repository interface {
    GetAll() ([]Task, error)     // Retrieve all tasks
    Save([]Task) error           // Persist tasks to storage
}
```

This interface-based approach provides several benefits:
- **Decoupling**: Business logic doesn't depend on specific storage implementations
- **Testability**: Easy to mock the repository for unit testing
- **Extensibility**: New storage backends (database, API, etc.) can implement this interface without changing business logic
- **Flexibility**: Multiple implementations can coexist for different use cases

#### `service.go` - Business Logic
Contains the `Service` struct that implements business logic operations:
- **Dependency Injection**: Accepts a `Repository` instance in the constructor
- **Validation**: Ensures task titles are non-empty
- **ID Generation**: Automatically assigns sequential IDs to new tasks
- **State Management**: Handles the complete task creation workflow

The service encapsulates all task-related business rules and operations.

### 3. **Storage Layer** (`internal/storage/`)

#### `json.go` - JSON Storage Implementation
Implements the `Repository` interface for JSON file-based persistence:
- **GetAll()**: Reads and deserializes tasks from `data/tasks.json`
- **Save()**: Serializes and writes tasks with formatted indentation for readability
- **Error Handling**: Gracefully handles missing files during read operations

This implementation demonstrates how to satisfy an interface contract for real-world data persistence.

## How It Works

### Current Functionality

The application currently supports the following command:

```bash
go run cmd/task/main.go add "Your task title"
```

**Workflow:**
1. Parse the "add" command and task title from arguments
2. Initialize JSON storage pointing to `data/tasks.json`
3. Create a service with the storage implementation
4. Call the service's `Add()` method to:
   - Validate the task title
   - Load existing tasks from storage
   - Generate the next sequential ID
   - Create a new task with `Done: false` status
   - Persist all tasks back to storage
5. Display a success message to the user

### Example Usage

```bash
# Add a new task
go run cmd/task/main.go add "Buy groceries"
# Output: Task added successfully

# Add another task
go run cmd/task/main.go add "Complete project"
# Output: Task added successfully
```

## Design Patterns

This project demonstrates several important Go design patterns:

### Dependency Injection
The `Service` struct receives its dependencies (Repository) through the constructor, rather than creating them internally. This makes the code testable and flexible.

### Interface Segregation
The `Repository` interface is small and focused, defining only the operations needed for task persistence. This makes it easy to implement and understand.

### Single Responsibility
Each package has a clear, focused purpose:
- `main`: Orchestration
- `task`: Business logic and domain models
- `storage`: Data persistence

## Package Visibility

The `internal/` directory is crucial for this project's structure:
- Packages within `internal/` **cannot be imported** by external modules
- This enforces encapsulation and prevents external code from depending on internal implementation details
- Only the CLI tool in `cmd/` is meant to be used externally

## Future Enhancements

The architecture supports easy expansion:
- **List tasks**: Add a `GetAll()` method to the service
- **Mark complete**: Add a `Complete(id)` method with repository operations
- **Delete tasks**: Add a `Delete(id)` method
- **Database storage**: Implement the Repository interface for SQL/NoSQL databases
- **REST API**: Create an HTTP server wrapper around the service
- **Configuration**: Add CLI flags for output format, storage location, etc.
