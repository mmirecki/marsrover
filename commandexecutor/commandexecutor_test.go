package commandexecutor_test

import (
	"github.com/mmirecki/marsrover/commandexecutor"
	"github.com/mmirecki/marsrover/rover"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type roverMock struct {
	forward, back, left, right int
}

func (r *roverMock) MoveForward() {
	r.forward++
}

func (r *roverMock) MoveBack() {
	r.back++
}

func (r *roverMock) TurnLeft() {
	r.left++
}

func (r *roverMock) TurnRight() {
	r.right++
}

func (r *roverMock) GetPosition() string {
	return ""
}

var _ = Describe("Commandexecutor", func() {
	var r *roverMock
	var executor commandexecutor.CommandExecutor

	BeforeEach(func() {
		r = &roverMock{}
		var _ rover.RoverInterface = r // ensure mock implements RoverInerface
		executor = commandexecutor.NewCommandExecutor(r)
	})

	It("verifies engine processes the forward command correctly", func() {
		err := executor.ExecuteCommand("F")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.forward).To(Equal(1))
	})

	It("verifies engine processes the back command correctly", func() {
		err := executor.ExecuteCommand("B")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.back).To(Equal(1))
	})

	It("verifies engine processes the left command correctly", func() {
		err := executor.ExecuteCommand("L")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.left).To(Equal(1))
	})

	It("verifies engine processes the right command correctly", func() {
		err := executor.ExecuteCommand("R")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.right).To(Equal(1))
	})

	It("verifies engine processes multiple commands correctly", func() {
		err := executor.ExecuteCommand("FFRRBLBFRF")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.forward).To(Equal(4))
		Expect(r.back).To(Equal(2))
		Expect(r.right).To(Equal(3))
		Expect(r.left).To(Equal(1))
	})

	It("verifies engine checks for invalid commands", func() {
		err := executor.ExecuteCommand("FFRRBLxBFRF")
		Expect(err).To(HaveOccurred())

		By("checkng the rover has not moved")
		Expect(r.forward).To(Equal(0))
		Expect(r.back).To(Equal(0))
		Expect(r.right).To(Equal(0))
		Expect(r.left).To(Equal(0))
	})
})
