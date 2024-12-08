package main

func add() {
	newTask := Task{
		Id:   getFreeId(),
		Task: fAdd,
		Done: false,
	}

	tasks = append(tasks, newTask)
	updateJson = true
}
