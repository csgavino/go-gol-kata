package csgavino_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/csgavino"
	"github.com/onsi/gomega"
)

var _ = Describe("Gol", func() {
	Context("Neighbours", func() {
		It("returns the number of neighbours", func() {
			gol := csgavino.NewGol(
				csgavino.Cell{X: 0, Y: 0, Alive: true},
				csgavino.Cell{X: 0, Y: 1, Alive: true},
				csgavino.Cell{X: 0, Y: 2, Alive: true},
				//
				csgavino.Cell{X: 1, Y: 0, Alive: true},
				csgavino.Cell{X: 1, Y: 2, Alive: true},
				//
				csgavino.Cell{X: 2, Y: 0, Alive: true},
				csgavino.Cell{X: 2, Y: 1, Alive: true},
				csgavino.Cell{X: 2, Y: 2, Alive: true},
			)

			neighbours := gol.Neighbours(1, 1)

			gomega.Expect(neighbours).To(gomega.Equal(8))
		})
	})

	Context("At", func() {
		It("returns the cell at given coordinate", func() {
			gol := csgavino.NewGol(
				csgavino.Cell{X: 0, Y: 0},
				csgavino.Cell{X: 0, Y: 1},
				csgavino.Cell{X: 0, Y: 2},
			)

			cell := gol.At(0, 1)

			gomega.Expect(cell).ToNot(gomega.BeNil())
			gomega.Expect(*cell).To(gomega.Equal(csgavino.Cell{X: 0, Y: 1}))
		})

		It("returns nil if no cell at given coordinate", func() {
			gol := csgavino.NewGol(
				csgavino.Cell{X: 0, Y: 0},
				csgavino.Cell{X: 0, Y: 1},
				csgavino.Cell{X: 0, Y: 2},
			)

			cell := gol.At(1, 1)

			gomega.Expect(cell).To(gomega.BeNil())
		})
	})
})
