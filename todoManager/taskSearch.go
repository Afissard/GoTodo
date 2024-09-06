package todomanager

import "fmt"

func (t TodoList) Search(nameSearched string) (int, error) {
	for i := 0; i < len(t.TaskList); i++ {
		if t.TaskList[i].Name == nameSearched {
			return i, nil
		}
	}
	return -1, fmt.Errorf("name not found")
}
