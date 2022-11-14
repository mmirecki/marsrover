package rover

const (
	North = iota
	West
	South
	East
)

type RoverInterface interface {
	MoveForward()
	MoveBack()
	TurnLeft()
	TurnRight()
	GetPosition() string
}
