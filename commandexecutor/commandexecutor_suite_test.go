package commandexecutor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCommandexecutor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commandexecutor Suite")
}
