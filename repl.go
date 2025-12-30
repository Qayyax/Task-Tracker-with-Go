package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Qayyax/Task-Tracker-with-Go/types"
)

func repl() {
	cliAscii := `
 /$$$$$$$$                  /$$               /$$$$$$  /$$       /$$$$$$
|__  $$__/                 | $$              /$$__  $$| $$      |_  $$_/
   | $$  /$$$$$$   /$$$$$$$| $$   /$$       | $$  \__/| $$        | $$  
   | $$ |____  $$ /$$_____/| $$  /$$//$$$$$$| $$      | $$        | $$  
   | $$  /$$$$$$$|  $$$$$$ | $$$$$$/|______/| $$      | $$        | $$  
   | $$ /$$__  $$ \____  $$| $$_  $$        | $$    $$| $$        | $$  
   | $$|  $$$$$$$ /$$$$$$$/| $$ \  $$       |  $$$$$$/| $$$$$$$$ /$$$$$$
   |__/ \_______/|_______/ |__/  \__/        \______/ |________/|______/

	`
	fmt.Println(cliAscii)
	info := `
	Welcome to Task-cli. To get started use the following command

	'help' - To list all the commands
	'exit' - To exit 
	`

	commands := `
	'add'	 - add [name] [description (optional)]
				This adds a new task to your list. You can pass a description if you want

	'update'	 - update [task-id] [name] [description (optional)]
				This updates an existing task name by id. You can update the desctiption as well 

	'delete'  - delete [task-id]
				This deletes an existing task by id.  

	'mark-todo' 			 - mark-todo [task-id]
	'mark-in-progress' - mark-in-progress [task-id]
	'mark-done' 			 - mark-done [task-id]

	'list'						 - list 
				This lists all the tasks you have
	 			list has optional arguments
				list todo
				list in-progress
				list done
	`

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(info)
	for {
		fmt.Print("task-cli ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			} else {
				fmt.Println("\nstdin closed (EOF). Bye!")
			}
			break
		}

		input := strings.TrimSpace(scanner.Text())
		args, err := splitArgs(input)
		if err != nil {
			fmt.Println("Invalid argument")
			continue
		}

		if input == "" {
			continue
		}

		command := args[0]
		if command == "exit" {
			fmt.Println("See you again")
			return
		}

		if command == "help" {
			fmt.Println(commands)
			continue
		}

		// if input has more than one argument,
		// run based on that
		// use the args
		// ADD command
		if command == "add" {
			if len(args) > 1 {
				name := args[1]
				description := ""
				if len(args) > 2 {
					fmt.Println(args[1])
					fmt.Println(args[2])
					description = args[2]
				}
				task := types.NewTask(name, description)
				types.AddToTasks(*task)
			} else {
				fmt.Println("Invalid use of command use 'help'")
				continue
			}
		}

		// UPDATE command
		if command == "update" {
			if len(args) > 2 {
				taskId := args[1]
				name := args[2]
				description := ""
				if len(args) > 3 {
					description = args[3]
				}
				taskIdNum, err := strconv.Atoi(taskId)

				if err != nil {
					fmt.Println("Input valid id number", err)
					continue
				}
				types.UpdateTask(taskIdNum, name, description)
			} else {
				fmt.Println("Invalid use of command use 'help' for more info")
			}

		}
		// DELETE command
		if command == "delete" {
			if len(args) > 1 {
				taskId := args[1]
				taskIdNum, err := strconv.Atoi(taskId)
				if err != nil {
					fmt.Println("Input valid id number", err)
					continue
				}
				types.DeleteTask(taskIdNum)
			} else {
				fmt.Println("Invalid use of command use 'help' for more info")
			}
		}
		// MARK-TODO command
		// MARK-IN-PROGRESS command
		// MARK-DONE command
	}
}

func splitArgs(input string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ' '
	r.LazyQuotes = true

	return r.Read()
}
