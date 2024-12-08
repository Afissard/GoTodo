package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
)

func init() {
	// command line flags

	flag.BoolVar(&fShow, "show", false, "show all tasks")
	flag.StringVar(&fAdd, "add", "", "add a task")
	flag.IntVar(&fDone, "done", -1, "mark a task as done")
	flag.IntVar(&fRemove, "remove", -1, "remove a task")
	flag.BoolVar(&fClear, "clear", false, "clear all tasks")

	flag.Parse()
}

func init() {
	// json creation / loading
	if _, err := os.Stat(jsonDataFile); errors.Is(err, os.ErrNotExist) {
		// create an empty json file
		file, err := os.Create(jsonDataFile)
		if err != nil {
			log.Fatal(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
		writeJson()

	} else {
		jsonContent, err := os.ReadFile(jsonDataFile)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(string(jsonContent))

		if len(jsonContent) != 0 {
			err = json.Unmarshal(jsonContent, &tasks)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	if fClear {
		clearTask()
	}

	if fRemove != -1 {
		remove()
	}

	if fAdd != "" {
		add()
	}

	if fDone != -1 {
		done()
	}

	if fShow {
		show()
	}

	if updateJson {
		writeJson()
	}
}
