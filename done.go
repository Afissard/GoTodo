package main

import (
	"fmt"
	"log"
)

func done() {
	foundId := false
	for i := 0; i < len(tasks); i++ {
		if i+1 == fDone {
			tasks[i].Done = !tasks[i].Done
			foundId = true
		}
	}

	if !foundId {
		log.Println(fmt.Errorf("id %d not found", fDone))
	} else {
		updateJson = true
	}
}
