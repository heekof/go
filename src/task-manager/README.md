# Task Manager

A simple command-line task management application written in Go.

## Overview

This task manager allows you to add, list, and complete tasks through a command-line interface. It's designed to be lightweight and easy to use for basic task tracking.

## Features

- **Add tasks**: Create new tasks with descriptions
- **List tasks**: View all tasks and their status
- **Complete tasks**: Mark tasks as completed by ID

## Installation

1. Make sure you have Go 1.23.4 or later installed
2. Clone or download this repository
3. Navigate to the task-manager directory
4. Build the application:
   ```bash
   go build -o task main.go
   ```

## Usage

### Basic Commands

```bash
# Add a new task
./task add "Complete the project documentation"

# List all tasks
./task list

# Complete a task by ID
./task complete 1
```

### Command Reference

- `task add <description>` - Add a new task with the given description
- `task list` - Display all tasks with their IDs and completion status
- `task complete <id>` - Mark the task with the given ID as completed

## Project Structure

```
task-manager/
├── main.go    # Main application code
├── go.mod     # Go module definition
└── README.md  # This file
```

## Code Structure

The application consists of:

- **Task struct**: Defines the task data structure with ID, Description, and Completed fields
- **main()**: Handles command-line argument parsing and routing
- **addTask()**: Handles adding new tasks
- **listTasks()**: Handles displaying all tasks
- **completeTask()**: Handles marking tasks as completed

## Current Status

⚠️ **Note**: This is currently a basic implementation with stub functions. The following functionality needs to be implemented:

- Task persistence (saving/loading tasks from file or database)
- Actual task management logic
- Task ID generation and validation
- Error handling for invalid task IDs

## Development

To run the application in development mode:

```bash
go run main.go <command> [args]
```

To run tests (when implemented):

```bash
go test
```

## Requirements

- Go 1.23.4 or later

## License

This project is open source and available under standard licensing terms.