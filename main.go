package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Qayyax/Task-Tracker-with-Go/types"
)

func main() {
	fmt.Println("Testing the Task struct")
	// fmt.Println("Current time is: ", time.Now())
	task1 := types.NewTask("First task", "Just to see if this works")
	fmt.Println("Name from new Task", task1.Name)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is your name? ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("You entered: %s\n", input)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	readFile()
}

// TODO:
// - [] How to edit files in goLang
