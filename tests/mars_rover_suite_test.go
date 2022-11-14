package tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMarsRover(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MarsRover Suite")
}
