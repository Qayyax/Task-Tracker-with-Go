package main

import (
	"bufio"
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

func printLastNlines(lines []string, num int) []string {
	var printLastNlines []string
	for i := len(lines) - num; i < len(lines); i++ {
		printLastNlines = append(printLastNlines, lines[i])
	}
	return printLastNlines
}

func printLogFile() {
	filepath := "log.txt"
	// we want to open this file to read line by line
	logFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer logFile.Close()
	var lines []string
	scanner := bufio.NewScanner(logFile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// print the last 10 lines of the file
	printLastNlines := printLastNlines(lines, 10)
	for i, line := range printLastNlines {
		fmt.Println("line number:", ":", i, line)
		fmt.Println("________")
	}
}
