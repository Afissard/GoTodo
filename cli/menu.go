package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetMenuChoice(menuOptions []string, timeFailed int) (choice int) {
	for i := 0; i < len(menuOptions); i++ {
		fmt.Printf("%d. %s", i+1, menuOptions[i])
	}
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	textInput, _ := reader.ReadString('\n')
	
	choice, _ = strconv.Atoi(textInput)
	if choice < 1 && choice > len(menuOptions) {
		if timeFailed +1 >= 3 {
			fmt.Println("3rd wrong input, exiting")
			os.Exit(0)
		}
		return GetMenuChoice(menuOptions, timeFailed+1)
	}
	return choice - 1
}
