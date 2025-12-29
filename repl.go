package main

import "fmt"

// make the following functions
// add - one or 2 positional argument  (add name and description)
// update - 2 or 3 positional argument (update id name description)
// delete - 1 positional argument (delete id)
// mark-todo - 1 positional argument (mark-todo id)
// mark-in-progress - 1 positional argument (mark-in-progress id)
// mark-done - 1 positional argument (mark-done id)
// list - no positional argument (list)
// positional arguments for list
// list todo
// list in-progress
// list done
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
}
