package main

import (
	"fmt"
	"os"
)

func readFile() {
	filepath := "text.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println("====This is the start of the text file====")
	fmt.Println(string(data))
	fmt.Println("====This is the end of the text file====")
}
