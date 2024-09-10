package cli

type menuOption struct {
	name string
	id   int
}

type menu []menuOption

var showTasks = menuOption{
	name: "Show task list",
	id:   1,
}

var addTask = menuOption{
	name: "Add a task",
	id:   2,
}

var removeTask = menuOption{
	name: "Remove a task",
	id:   3,
}

var quit = menuOption{
	name: "Quit",
	id:   4,
}

var mainMenuOptions = menu{showTasks, addTask, removeTask, quit}
