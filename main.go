package main

import (
	"bufio"
	"fmt"
	"github.com/mmirecki/marsrover/commandexecutor"
	"github.com/mmirecki/marsrover/rover"
	"os"
	"strings"
)

func main() {
	marsRover := rover.NewRover(0, 0, rover.North)
	commandExecutor := commandexecutor.NewCommandExecutor(marsRover)
	reader := bufio.NewReader(os.Stdin)
	congrats := []string{"Congratulations!", "Good Job!", "Well done!"}

	fmt.Println("Enter command for the Mars rover (FBRL)")
	fmt.Println("('q' to exit)")
	quitReceived := false
	for i:=0; !quitReceived; i=(i+1)%len(congrats) {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Trim(userInput, "\n")
		if userInput == "" {
			continue
		}
		switch userInput {
		case "q":
			fmt.Println("Mission accomplished! The Mars Rover has reached the desired destination!")
			quitReceived = true
		default:
			fmt.Printf("Executing command: %s\n", userInput)
			err := commandExecutor.ExecuteCommand(userInput)
			if err != nil {
				fmt.Printf("Invalid command: %s\n", err)
				continue
			}
			fmt.Printf("%s Command executed successfully\n", congrats[i] )
			fmt.Printf("The new Mars Rover position is %s\n\n", marsRover.GetPosition())
		}
	}
}
