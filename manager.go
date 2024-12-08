package main

import "sort"

func sortIdTask() {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})
}

func getFreeId() int {
	if len(tasks) == 0 {
		return 1
	}
	sortIdTask()
	return tasks[len(tasks)-1].Id + 1
}

func countTaskDone() int {
	count := 0
	for _, task := range tasks {
		if task.Done == true {
			count++
		}
	}
	return count
}
