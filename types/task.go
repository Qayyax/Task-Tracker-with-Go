package types

import "time"

// - `id`: A unique identifier for the task
// - `description`: A short description of the task
// - `status`: The status of the task (`todo`, `in-progress`, `done`)
// - `createdAt`: The date and time when the task was created
// - `updatedAt`: The date and time when the task was last updated

type Task struct {
	Id          int
	Name        string
	Description string
	Status      Status // [0] todo, [1] in-progress, [2] done
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

// TODO:
// - [] methods for Task
// 	- [] Add new task
// 	- [] Update an old task based on Id
// 	- [] Delete a task based on id
// 	- [] List all tasks
// 	- [] List all tasks that are not done [0]
// 	- [] List all tasks that are in progress [1]
// 	- [] List all tasks that are done [2]
// - [] cli repl question that asks for name of new task and the description

// This is not a method
func NewTask(name string, description string) *Task {
	// should check for last id in the json file

	task := Task{
		Id:          1, // change this to check for id
		Name:        name,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
	}
	return &task
}
