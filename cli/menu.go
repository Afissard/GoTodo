package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getMenuChoice(options menu, timeFailed int) (choice int) {
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d. %s\n", options[i].id, options[i].name)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your choice: ")
	textInput, _ := reader.ReadString('\n')

	choice, _ = strconv.Atoi(textInput)
	if choice < 1 || choice > len(options) { // TODO : validity test not up to date with new menu système
		timeFailed++
		fmt.Printf("invalid input : %d\n", choice)
		if timeFailed >= 3 {
			fmt.Println("3rd invalid input, exiting")
			os.Exit(0)
		}
		return getMenuChoice(options, timeFailed+1)
	}
	return choice
}
