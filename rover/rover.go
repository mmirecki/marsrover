package rover

type RoverInterface interface {
	MoveForward()
	MoveBack()
	TurnLeft()
	TurnRight()
	GetPosition() string
}