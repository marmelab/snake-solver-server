package computer_test

import (
    . "github.com/marmelab/snake-solver-server/lib"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

const width, height = 5, 5

func TestComputer(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Computer")
}

var _ = Describe("Computer", func() {
    It("should be return snake head", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }
        Expect(GetSnakeHead(snake)).To(Equal([2]int{0, 2}))
    })

    It("should check empty cell", func() {
        grid := [width][height]int{
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }

        Expect(IsEmptyCell(grid, [2]int{0, 2})).To(Equal(false))
        Expect(IsEmptyCell(grid, [2]int{2, 2})).To(Equal(true))
    })

    It("it should check outside bounding box", func() {
        grid := [width][height]int{
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }

        Expect(IsOutsideBoundingBox([2]int{-1, 0}, grid)).To(Equal(true))
        Expect(IsOutsideBoundingBox([2]int{0, 0}, grid)).To(Equal(false))
    })
})
