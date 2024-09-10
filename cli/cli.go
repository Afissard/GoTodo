package cli

const (
	// main menu
	mainMenuOptions := [
		"Show task list", 
		"Add a task",
		"Remove a task",
		"Quit"
		]

	showTasks int = iota
	addTask
	removeTask
	quit
)