package csgavino_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/csgavino"
)

var _ = Describe("Cell", func() {
	Describe("Alive", func() {
		Context("with less than 2 neighbours", func() {
			It("dies", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      true,
					Neighbours: 1,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeFalse())
			})
		})

		Context("with 2 or 3 neighbours", func() {
			It("lives", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      true,
					Neighbours: 2,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeTrue())
			})

			It("lives", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      true,
					Neighbours: 3,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeTrue())
			})
		})

		Context("with more than 3 neighbours", func() {
			It("dies", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      true,
					Neighbours: 4,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeFalse())
			})
		})

	})

	Describe("Dead", func() {
		Context("with 2 neighbours", func() {
			It("stays dead", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      false,
					Neighbours: 2,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeFalse())
			})
		})

		Context("with 3 neighbours", func() {
			It("lives", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      false,
					Neighbours: 3,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeTrue())
			})
		})

		Context("with 4 neighbours", func() {
			It("stays dead", func() {
				cell := csgavino.Cell{
					X:          1,
					Y:          1,
					Alive:      false,
					Neighbours: 4,
				}

				cell.Tick()

				gomega.Expect(cell.Alive).Should(gomega.BeFalse())
			})
		})

	})

})
