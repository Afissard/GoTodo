package todomanager

import "fmt"

func (t TaskList) Search(nameSearched string) (int, error) {
	for i := 0; i < len(t.Tasks); i++ {
		if t.Tasks[i].Name == nameSearched {
			return i, nil
		}
	}
	return -1, fmt.Errorf("name not found")
}
