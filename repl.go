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
}
