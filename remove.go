package main

import (
	"fmt"
	"log"
)

func remove() {
	foundId := false
	for i := 0; i < len(tasks); i++ {
		if i+1 == fRemove {
			foundId = true
		}
	}

	if !foundId {
		log.Println(fmt.Errorf("id %d not found", fRemove))
	} else {
		tasks = append(tasks[:fRemove-1], tasks[fRemove:]...)
		updateJson = true
	}
}
