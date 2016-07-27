package computer_test

import (
    . "github.com/marmelab/snake-solver-server/lib"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

const width, height = 5, 5
const up, right, down, left = 0, 1, 2, 3

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

    It("should check outside bounding box", func() {
        grid := [width][height]int{
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }

        Expect(IsOutsideBoundingBox([2]int{-1, 0}, grid)).To(Equal(true))
        Expect(IsOutsideBoundingBox([2]int{0, 5}, grid)).To(Equal(true))
        Expect(IsOutsideBoundingBox([2]int{0, 0}, grid)).To(Equal(false))
    })

    It("should return adjacent position", func() {
        Expect(GetAdjacentPosition([2]int{1, 0}, up)).To(Equal([2]int{0, 0}))
        Expect(GetAdjacentPosition([2]int{0, 0}, right)).To(Equal([2]int{0, 1}))
        Expect(GetAdjacentPosition([2]int{0, 0}, down)).To(Equal([2]int{1, 0}))
        Expect(GetAdjacentPosition([2]int{0, 1}, left)).To(Equal([2]int{0, 0}))
    })

    It("should return possible moves", func() {
        grid := [width][height]int{
            {1, 1, 1, 0, 0},
            {0, 0, 1, 0, 0},
            {2, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
        }

        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
            {1, 2},
            {2, 2},
            {2, 1},
        }

        Expect(GetPossibleMoves(grid, snake)).To(Equal([]int{0, 2, 3}))
    })

    It("should move snake", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        newSnake := MoveSnake(snake, []int{right, right, down})

        Expect(newSnake).To(Equal([][2]int{
            {0, 3},
            {0, 4},
            {1, 4},
        }))
    })

    It("should initialize grid", func() {
        snake := [][2]int{
            {2, 0},
            {2, 1},
            {2, 2},
        }

        apple := [2]int{4, 4}

        grid := InitializeGrid(snake, apple)
        Expect(grid).To(Equal([width][height]int{
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }))
    })

    It("should find path", func() {
        grid := [width][height]int{
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 2, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
        }

        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        apple := [2]int{2, 2}

        path := GetPath(grid, snake, apple)
        Expect(path).To(Equal([]int{down, down}))
    })
})
