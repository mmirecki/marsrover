package tests_test

import (
	"github.com/mmirecki/marsrover/commandexecutor"
	"github.com/mmirecki/marsrover/rover"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MarsRover", func() {
	var r rover.RoverInterface // starting postion: (0, 0) NORTH
	var executor commandexecutor.CommandExecutor

	BeforeEach(func() {
		r = rover.NewRover(0, 0, rover.North)
		executor = commandexecutor.NewCommandExecutor(r)
	})

	It("verifies engine processes the forward command correctly", func() {
		err := executor.ExecuteCommand("F")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 1) NORTH"))
	})

	It("verifies engine processes the back command correctly", func() {
		err := executor.ExecuteCommand("B")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, -1) NORTH"))
	})

	It("verifies engine processes the turn left command correctly", func() {
		err := executor.ExecuteCommand("L")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 0) WEST"))
	})

	It("verifies engine processes the turn right command correctly", func() {
		err := executor.ExecuteCommand("R")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 0) EAST"))
	})

	It("verifies engine processes multiple forward commands correctly", func() {
		err := executor.ExecuteCommand("FFFFFFFF")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 8) NORTH"))
	})

	It("verifies engine processes multiple back commands correctly", func() {
		err := executor.ExecuteCommand("BBBBB")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, -5) NORTH"))
	})

	It("verifies engine processes multiple turn right commands correctly", func() {
		err := executor.ExecuteCommand("RRRRRRR")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 0) WEST"))
	})

	It("verifies engine processes multiple turn left commands correctly", func() {
		err := executor.ExecuteCommand("LLLLL")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 0) WEST"))
	})

	It("verifies engine processes multiple forward and turn commands correctly", func() {
		err := executor.ExecuteCommand("FRFLFRFLFRFLFRFLFRFLFRFLFRFLFRFL")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(8, 8) NORTH"))
	})

	It("verifies engine processes multiple forward, back and turn commands correctly", func() {
		err := executor.ExecuteCommand("FLBRFLBRFLBRFLBRFLBRFLBRFLBRFLBR")
		Expect(err).NotTo(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(8, 8) NORTH"))
	})

	It("verifies engine detects invalid command without moving rover", func() {
		err := executor.ExecuteCommand("FLBRFLBRFLBRFLBRfFLBRFLBRFLBRFLBR")
		Expect(err).To(HaveOccurred())
		Expect(r.GetPosition()).To(Equal("(0, 0) NORTH"))
	})
})
