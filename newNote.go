package main

import (
	"bufio"
	"fmt"
	"os"
)

func newTask() {
	filepath := "task.txt"
	taskFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", err)
		// we would create a new file
		newTaskFile, err := os.Create(filepath)
		if err != nil {
			panic(err)
		}
		defer newTaskFile.Close()

	}
	defer taskFile.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your new task here")
	scanner.Scan()
	input := scanner.Text()

	_, err = taskFile.Write([]byte(input))

	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	fmt.Println("====This is the start of the text file====")
	fmt.Println(string(data))
	fmt.Println("====This is the end of the text file====")
}
