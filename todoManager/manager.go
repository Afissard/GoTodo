package todomanager

type task struct {
	Name        string
	Description string
	Priority    int // the higher the value, the higher the priority
	Status      int
}

type TaskList struct {
	TaskList []task
}

const ( // Status
	NotStarted int = iota
	Started
	Finished
)
