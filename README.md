# Task Tracker with Go

A simple REPL-based CLI to manage tasks (add, update, delete, list, mark status). Tasks are persisted to `tasks.json` in the working directory.

This project was my introduction to Go; I learned the language syntax while building it.

## Features

- Add tasks with an optional description
- Update task name and description by id
- Mark tasks as todo, in-progress, or done
- List all tasks or filter by status
- Persist tasks to a local JSON file

## Requirements

- Go 1.25+ (see `go.mod`)

## Run without building

```bash
go run .
```

## Build

```bash
go build -o task-cli .
```

## Run the binary

```bash
./task-cli        # macOS/Linux
task-cli.exe      # Windows
```

Note: Run the binary from the project root (or a directory containing `tasks.json`), because storage uses a relative path.

## Usage

The CLI opens an interactive prompt. You can use these commands:

```text
help
exit

add "name" "description"
update 1 "name" "description"
delete 1

mark-todo 1
mark-in-progress 1
mark-done 1

list
list todo
list in-progress
list done
```

Tips:
- Wrap multi-word names or descriptions in quotes.
- `description` is optional for `add` and `update`.

## Data storage

- Tasks are stored in `tasks.json` in the current working directory.
- The file is created on first run if it does not exist.
