package types

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// - `id`: A unique identifier for the task
// - `description`: A short description of the task
// - `status`: The status of the task (`todo`, `in-progress`, `done`)
// - `createdAt`: The date and time when the task was created
// - `updatedAt`: The date and time when the task was last updated

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// [0] todo, [1] in-progress, [2] done
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	isTaskFile := IsTaskFileExist()
	if isTaskFile != true {
		panic("Didn't create a new file")
	}
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

func IsTaskFileExist() bool {
	filepath := "tasks.json"
	data, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Tasks.json does not exist")

		newTaskFile, err := os.Create(filepath)
		if err != nil {
			panic(err)
		}
		defer newTaskFile.Close()
	}
	defer data.Close()
	return true
}

func checkLastTaskId() int {
	IsTaskFileExist()
	filepath := "tasks.json"
	var tasks []Task

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	lastTask := tasks[len(tasks)-1]
	return lastTask.Id
}
