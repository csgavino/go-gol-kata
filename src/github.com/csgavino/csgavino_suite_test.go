package csgavino_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCsgavino(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gol Suite")
}
