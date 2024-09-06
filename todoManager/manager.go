package todomanager

type task struct {
	Name        string
	Description string
	Priority    int // the higher the value, the higher the priority
	Status      int
}

type TodoList struct {
	TaskList []task
}

const ( // Status
	NotStarted int = iota
	Started
	Finished
)
