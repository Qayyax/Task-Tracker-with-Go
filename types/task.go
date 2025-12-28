package types

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
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
// 	- [x] Add new task
// 	- [x] Update an old task based on Id
// 	- [] Delete a task based on id
// 	- [] List all tasks
// 	- [] List all tasks that are not done [0]
// 	- [] List all tasks that are in progress [1]
// 	- [] List all tasks that are done [2]
// - [] cli repl question that asks for name of new task and the description

func NewTask(name string, description string) *Task {
	isTaskFile := IsTaskFileExist()
	if isTaskFile != true {
		panic("Didn't create a new file")
	}
	lastId := checkLastTaskId() + 1
	task := Task{
		Id:          lastId,
		Name:        name,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
	if len(tasks) < 1 {
		return 0
	}

	lastTask := tasks[len(tasks)-1]
	return lastTask.Id
}

func AddToTasks(task Task) {
	filepath := "tasks.json"
	var tasks []Task

	content, err := os.ReadFile(filepath)
	if err == nil && len(content) > 0 {
		json.Unmarshal(content, &tasks)
	}
	tasks = append(tasks, task)

	updateTasksFile, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Printf("Error encoding JSON to %s\nERROR: %v", filepath, err)
		return
	}

	err = os.WriteFile(filepath, updateTasksFile, 0644)
	if err != nil {
		fmt.Println("Error writing json to file:", err)
		return
	}
	fmt.Println("Added new task to tasks.json")
}

// get task index by id then update it
func GetTaskIndexById(id int) int {
	filepath := "tasks.json"
	var tasks []Task

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Didn't decode tasks", err)
	}
	fmt.Println(tasks)
	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.Id == id
	})
	return index
}

// update the task based on the the index
// if only one string is passed, it would be just the name,
// if a second string is passed it would be the description
func UpdateTask(id int, name string, description string) {
	index := GetTaskIndexById(id)
	if index < 0 {
		// The id can't be found in the tasks
		return
	}
	filepath := "tasks.json"
	var tasks []Task

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		fmt.Println("Didn't decode tasks", err)
	}
	oldName := tasks[index].Name
	tasks[index].UpdatedAt = time.Now()
	tasks[index].Name = name
	if description != "" {
		tasks[index].Description = description
	}
	for index, task := range tasks {
		fmt.Printf("%v: %+v\n", index, task)
	}

	updateTasksFile, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Printf("Error encoding JSON to %s\nERROR: %v", filepath, err)
		return
	}

	err = os.WriteFile(filepath, updateTasksFile, 0644)
	if err != nil {
		fmt.Println("Error writing json to file:", err)
		return
	}
	fmt.Printf("Updated %v to %v", oldName, tasks[index].Name)
}

// Delete a task by ID
func DeleteTask(id int) {

}
