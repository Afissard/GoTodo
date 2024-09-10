package main

import (
	"TodoCLI/cli"
	todomanager "TodoCLI/todoManager"
)

func main() {
	todoList := todomanager.TaskList{}

	cli.MainMenu(todoList)
}
