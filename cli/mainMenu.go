package cli

import (
	todomanager "TodoCLI/todoManager"
	"fmt"
	"os"
)

func MainMenu(taskList todomanager.TaskList){
	switch getMenuChoice(mainMenuOptions, 0) {
	case mainMenuOptions[0].id:
		displayTasks(taskList)
	case mainMenuOptions[3].id:
		fmt.Println("Exit")
		os.Exit(0)
	default: fmt.Println("WIP or error")
	}
}