package main

import "fmt"

func show() {
	fmt.Println("Id\tTodo\tTask")
	for i := 0; i < len(tasks); i++ {
		fmt.Printf("%d\t%t\t%s\n", tasks[i].Id, tasks[i].Done, tasks[i].Task)
	}
	fmt.Printf("%d/%d done\n", countTaskDone(), len(tasks))
}
