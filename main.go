package main

import (
	"fmt"

	"github.com/Qayyax/Task-Tracker-with-Go/types"
)

func main() {
	fmt.Println("Testing the Task struct")
	// fmt.Println("Current time is: ", time.Now())
	task1 := types.NewTask("First task", "Just to see if this works")
	fmt.Println("Name", task1.Name)
	fmt.Println("Description", task1.Description)
	fmt.Println("CreatedAt", task1.CreatedAt)
}
