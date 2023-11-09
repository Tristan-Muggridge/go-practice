# Go Task Manager

## Description

A CLI task manager which allows for the creation, deletion, and listing of tasks.

## Setup

1. Clone the repository
2. Run
```bash
    cd task-manager/infra
    docker-compose -p go-task-manager up -d
```

## Usage

Run the following command to run an instance of the task manager
```bash
    go task-manager/src/cmd/run main.go
```

You'll be presented with a prompt to enter a command. The following commands are available:

1. Create Task
2. List Tasks
3. Edit Tasks
4. Delete Task
5. Exit

## Technology

- Storage: Postgres (Docker instances, one for test and one for dev)
- ORM: go-pg/v10
- Language: Go