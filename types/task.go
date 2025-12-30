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
// - [x] methods for Task
// 	- [x] Add new task
// 	- [x] Update an old task based on Id
// 	- [x] Delete a task based on id
// 	- [x] update status based on id
// 	- [x] List all tasks
// 	- [x] List all tasks that are not done [0]
// 	- [x] List all tasks that are in progress [1]
// 	- [x] List all tasks that are done [2]
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
	tasks, err := GetTasks()
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
		fmt.Println("Id does not exist")
		return
	}
	filepath := "tasks.json"
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println("Failed to fetch tasks", err)
		return
	}
	oldName := tasks[index].Name
	tasks[index].UpdatedAt = time.Now()
	tasks[index].Name = name
	if description != "" {
		tasks[index].Description = description
	}
	fmt.Println("=========")
	fmt.Println(tasks[index].Id)
	fmt.Println(tasks[index].Name)
	fmt.Println(tasks[index].Description)
	fmt.Println(tasks[index].UpdatedAt)
	fmt.Println("=========")

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
	fmt.Printf("Updated %v to %v\n", oldName, tasks[index].Name)
}

// Delete a task by ID
func DeleteTask(id int) {
	index := GetTaskIndexById(id)
	filepath := "tasks.json"
	if index < 0 {
		fmt.Println("Can't delete task with invalid id")
		fmt.Println("Please enter a valid id")
		return
	}
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	tasks = slices.Delete(tasks, index, index+1)
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
	fmt.Println("Successfully deleted task", id)
}

// Get all tasks from json
func GetTasks() ([]Task, error) {
	filepath := "tasks.json"
	var tasks []Task

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	return tasks, err
}

/*
Update status by id

it would take the id, and the option (0, 1, 2)
Would use switch case statement to switch task

id - Task id

option - [0] todo, [1] in-progress, [2] done
*/
func UpdateTaskStatus(id, option int) {
	index := GetTaskIndexById(id)
	if index < 0 {
		fmt.Println("ID does not exist")
		return
	}
	tasks, err := GetTasks()
	if err != nil {
		fmt.Printf("Failed to get all tasks")
		return
	}

	switch option {
	case 0:
		tasks[index].Status = Todo
		fmt.Printf("Task id: %v status updated to 'Todo'", id)
	case 1:
		tasks[index].Status = InProgress
		fmt.Printf("Task id: %v status updated to 'In progress'", id)
	case 2:
		tasks[index].Status = Done
		fmt.Printf("Task id: %v status updated to 'Done'", id)
	default:
		fmt.Println("Invalid option")
		return
	}

	filepath := "tasks.json"
	updatedTaskFile, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Failed to convert tasks to json", err)
		return
	}
	err = os.WriteFile(filepath, updatedTaskFile, 0644)
	if err != nil {
		fmt.Println("Error writing updated task json to file", err)
		return
	}
	fmt.Println("Successfully updated task status")
}

// Function to list all tasks
func ListAllTask() {
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println("Unable to get all tasks")
		return
	}

	if len(tasks) < 1 {
		fmt.Println("You have no tasks")
		return
	}

	for _, task := range tasks {
		fmt.Printf("You have %v tasks\n", len(tasks))
		fmt.Println("-------------")
		fmt.Printf("Id: %v\n", task.Id)
		fmt.Printf("Name: %v\n", task.Name)
		fmt.Printf("Description: %v\n", task.Description)
		switch task.Status {
		case Todo:
			fmt.Printf("Status: Todo\n")
		case InProgress:
			fmt.Printf("Status: In-progress\n")
		case Done:
			fmt.Printf("Status: Done\n")
		}
		fmt.Println("-------------")
	}
}

func ListAllTasksNotDone() {
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println("Unable to get all tasks")
		return
	}

	if len(tasks) < 1 {
		fmt.Println("You have no tasks")
		return
	}

	var tasksNotDone []Task

	for _, task := range tasks {
		if task.Status == Todo {
			tasksNotDone = append(tasksNotDone, task)
			fmt.Println("-------------")
			fmt.Printf("Id: %v\n", task.Id)
			fmt.Printf("Name: %v\n", task.Name)
			fmt.Printf("Description: %v\n", task.Description)
			fmt.Println("-------------")
		}
	}
	fmt.Printf("You have %v tasks not done\n", len(tasksNotDone))
}

func ListAllTasksInProgress() {
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println("Unable to get all tasks")
		return
	}

	if len(tasks) < 1 {
		fmt.Println("You have no tasks")
		return
	}

	var tasksInProgress []Task

	for _, task := range tasks {
		if task.Status == InProgress {
			tasksInProgress = append(tasksInProgress, task)
			fmt.Println("-------------")
			fmt.Printf("Id: %v\n", task.Id)
			fmt.Printf("Name: %v\n", task.Name)
			fmt.Printf("Description: %v\n", task.Description)
			fmt.Println("-------------")
		}
	}
	fmt.Printf("You have %v tasks in-progress\n", len(tasksInProgress))
}

func ListAllTasksDone() {
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println("Unable to get all tasks")
		return
	}

	if len(tasks) < 1 {
		fmt.Println("You have no tasks")
		return
	}

	var tasksCompleted []Task

	for _, task := range tasks {
		if task.Status == InProgress {
			tasksCompleted = append(tasksCompleted, task)
			fmt.Println("-------------")
			fmt.Printf("Id: %v\n", task.Id)
			fmt.Printf("Name: %v\n", task.Name)
			fmt.Printf("Description: %v\n", task.Description)
			fmt.Println("-------------")
		}
	}
	fmt.Printf("You have %v tasks marked as done\n", len(tasksCompleted))
}
