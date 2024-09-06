package todomanager

func New(name, description string, priority int) task {
	return task{
		Name:        name,
		Description: description,
		Priority:    priority,
	}
}
