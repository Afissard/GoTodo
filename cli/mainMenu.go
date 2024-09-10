package cli

import (
	todomanager "TodoCLI/todoManager"
	"fmt"
)

func MainMenu(taskList todomanager.TaskList){
	switch GetMenuChoice(mainMenuOptions, 0) {
	case showTasks:
		displayTasks(taskList)
	case quit:
		fmt.Println("Exit")
		Exit(0)
	}
}