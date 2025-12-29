package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		fmt.Println(input)
		if input == "" {
			continue
		}

		if input == "exit" {
			fmt.Println("See you again")
			return
		}

		if input == "help" {
			fmt.Println(commands)
			continue
		}
		// if input has more than one argument,
		// run based on that
	}
}
