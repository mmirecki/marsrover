package rover

import "fmt"

const (
	North = iota
	West
	South
	East
	forward = 1
	back    = -1
)

var (
	Directions       = []string{"NORTH", "WEST", "SOUTH", "EAST"}
	directionVectors = [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
)

type RoverInterface interface {
	MoveForward()
	MoveBack()
	TurnLeft()
	TurnRight()
	GetPosition() string
}

type Rover struct {
	x, y      int
	direction int
}

func NewRover(x, y int, direction int) *Rover {
	return &Rover{x: x, y: y, direction: direction}
}

func (rover *Rover) MoveForward() {
	rover.move(forward)
}

func (rover *Rover) MoveBack() {
	rover.move(back)
}

func (rover *Rover) TurnLeft() {
	rover.direction = (rover.direction + 5) % 4
}

func (rover *Rover) TurnRight() {
	rover.direction = (rover.direction + 3) % 4
}

func (rover *Rover) GetPosition() string {
	return fmt.Sprintf("(%d, %d) %s", rover.x, rover.y, Directions[rover.direction])
}

func (rover *Rover) move(direction int) {
	directionVector := directionVectors[rover.direction]
	rover.x += directionVector[0] * direction
	rover.y += directionVector[1] * direction
}
