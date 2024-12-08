package main

const (
	jsonDataFile = "gotodo.json"
)

var (
	// flags
	fShow   bool
	fAdd    string
	fDone   int
	fRemove int
	fClear  bool

	// app var
	tasks      []Task // maybe replace it with a map
	updateJson = false
)

type Task struct {
	Id   int
	Task string
	Done bool
}
