package rover_test

import (
	r "github.com/mmirecki/marsrover/rover"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rover", func() {
	var rover r.RoverInterface

	BeforeEach(func() {
		rover =	r.NewRover(3, 4, r.North)
	})

	It("verifies new rover reports correct position", func() {
		Expect(rover.GetPosition()).To(Equal("(3, 4) NORTH"))
	})

	It("verifies move forward works correctly", func() {
		rover.MoveForward()
		Expect(rover.GetPosition()).To(Equal("(3, 5) NORTH"))
	})

	It("verifies move backwards works correctly", func() {
		rover.MoveBack()
		Expect(rover.GetPosition()).To(Equal("(3, 3) NORTH"))
	})

	It("verifies turn right works correctly", func() {
		rover.TurnRight()
		Expect(rover.GetPosition()).To(Equal("(3, 4) EAST"))
	})

	It("verifies turn left works correctly", func() {
		rover.TurnLeft()
		Expect(rover.GetPosition()).To(Equal("(3, 4) WEST"))
	})

	It("verifies multiple commands together work correctly", func() {
		rover.TurnLeft()
		rover.TurnLeft()
		rover.MoveForward()
		rover.MoveForward()
		rover.TurnRight()
		rover.MoveBack()
		Expect(rover.GetPosition()).To(Equal("(4, 2) WEST"))
	})
})
