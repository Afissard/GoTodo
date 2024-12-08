package main

import (
	"encoding/json"
	"log"
	"os"
)

func writeJson() {
	file, err := os.OpenFile(jsonDataFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// clear the file
	err = file.Truncate(0)
	if err != nil {
		return
	}

	temp, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(temp)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Saved JSON to", jsonDataFile)
}

func clearJson() {
	file, err := os.OpenFile(jsonDataFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Truncate(0)
	if err != nil {
		return
	}
}
