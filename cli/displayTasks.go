package cli

import (
	todomanager "TodoCLI/todoManager"
	"fmt"
)

func DisplayTasks(taskList todomanager.TaskList) {
	output := "Priority\tStatus\tDescription\n"
	for i := 0; i < len(taskList.Tasks); i++ {
		currentTask := taskList.Tasks[i]
		output += string(currentTask.Priority) +"\t"+ currentTask.GetStatus() + "\t" + currentTask.Name + "\n"
	}
	fmt.Print(output)
}
