package commandexecutor

import (
	"errors"
	"fmt"
	"github.com/mmirecki/marsrover/rover"
	"strings"
	"sync"
)

const (
	AllowedCommands = "FBLR"
)

type CommandExecutor struct {
	sync.Mutex
	marsRover rover.RoverInterface
}

func NewCommandExecutor(marsRover rover.RoverInterface) CommandExecutor {
	return CommandExecutor{marsRover: marsRover}
}

func (executor *CommandExecutor) ExecuteCommand(commandString string) error {
	executor.Lock()
	defer executor.Unlock()

	commands := []string{}
	for _, command := range strings.Split(commandString, "") {
		if err := validateCommand(command); err != nil {
			return err
		}
		commands = append(commands, command)
	}
	for _, command := range commands {
		executor.processCommandString(command)
	}
	return nil
}

func (executor *CommandExecutor) processCommandString(commandString string) {

	switch commandString {
	case "F":
		executor.marsRover.MoveForward()
	case "B":
		executor.marsRover.MoveBack()
	case "L":
		executor.marsRover.TurnLeft()
	case "R":
		executor.marsRover.TurnRight()
	}
}

func validateCommand(command string) error {
	if !strings.Contains(AllowedCommands, command) {
		return errors.New(fmt.Sprintf("Invalid character %s in command string", command))
	}
	return nil
}