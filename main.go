package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Qayyax/Task-Tracker-with-Go/types"
)

func main() {
	fmt.Println("Testing the Task struct")
	// cli to ask for task name, task description
	// task1 := types.NewTask("First task", "Just to see if this works")
	// fmt.Println("Name from new Task", task1.Name)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is the name of your first task: ")
	scanner.Scan()
	taskName := scanner.Text()
	fmt.Printf("What is the description of your task?\n")
	scanner.Scan()
	taskDescription := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from input:", err)
		panic(err)
	}
	task1 := types.NewTask(taskName, taskDescription)
	fmt.Println("Name from new Task", task1.Name)
	fmt.Println("Description from new Task", task1.Description)
	types.AddToTasks(*task1)
}
