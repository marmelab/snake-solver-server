package computer

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

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
        Expect(getSnakeHead(snake)).To(Equal([2]int{0, 2}))
    })

    It("should check empty cell", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        Expect(isEmptyCell([2]int{0, 2}, snake)).To(Equal(false))
        Expect(isEmptyCell([2]int{2, 2}, snake)).To(Equal(true))
    })

    It("should check outside bounding box", func() {
        Expect(isOutsideBoundingBox(size{5, 5}, [2]int{-1, 0})).To(Equal(true))
        Expect(isOutsideBoundingBox(size{5, 5}, [2]int{0, 5})).To(Equal(true))
        Expect(isOutsideBoundingBox(size{5, 5}, [2]int{0, 0})).To(Equal(false))
    })

    It("should return adjacent position", func() {
        Expect(getAdjacentPosition([2]int{1, 0}, up)).To(Equal([2]int{0, 0}))
        Expect(getAdjacentPosition([2]int{0, 0}, right)).To(Equal([2]int{0, 1}))
        Expect(getAdjacentPosition([2]int{0, 0}, down)).To(Equal([2]int{1, 0}))
        Expect(getAdjacentPosition([2]int{0, 1}, left)).To(Equal([2]int{0, 0}))
    })

    /*
        {1, 1, 1, 0, 0},
        {0, 0, 1, 0, 0},
        {2, 1, 1, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
    */
    It("should return possible moves", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
            {1, 2},
            {2, 2},
            {2, 1},
        }

        Expect(getPossibleMoves(size{5, 5}, snake)).To(Equal([]int{0, 2, 3}))
    })

    /*
        {1, 1, 1, 0, 0},
        {1, 1, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
    */
    It("should check if snake has free space", func() {
        snake := [][2]int{
            {0, 2}, {0, 1}, {1, 1}, {1, 0}, {0, 0},
        }

        Expect(isSnakeHasFreeSpace(size{5, 5}, snake)).To(Equal(false))
    })

    It("should move snake", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        apple := [2]int{4, 0}

        newSnake := moveSnake(snake, apple, []int{right, right, down})

        Expect(newSnake).To(Equal([][2]int{
            {0, 3},
            {0, 4},
            {1, 4},
        }))
    })

    /*
        {1, 1, 1, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 2, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
    */
    It("should find path", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        apple := [2]int{2, 2}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:2]).To(Equal([]int{down, down}))
    })

    It("should find path (2)", func() {
        snake := [][2]int{
            {2, 1}, {1, 1}, {1, 2}, {0, 2}, {0, 3}, {1, 3}, {1, 4}, {2, 4}, {2, 3}, {3, 3}, {4, 3}, {4, 2},
        }

        apple := [2]int{4, 0}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:1][0]).To(Equal(left))
    })

    It("should eat last apple", func() {
        snake := [][2]int{
            {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {3, 1}, {2, 1}, {2, 2}, {3, 2}, {4, 2}, {4, 3},
             {4, 4}, {3, 4}, {3, 3}, {2, 3}, {2, 4}, {1, 4}, {0, 4}, {0, 3}, {1, 3}, {1, 2}, {0, 2}, {0, 1},
        }

        Expect(isLastMove(size{5, 5}, snake)).To(Equal(true))
        Expect(isLastMove(size{5, 5}, snake[:len(snake)-1])).To(Equal(false))

        path := GetPath(5, 5, snake, [2]int{1, 1})

        Expect(path[:1][0]).To(Equal(down))
    })

    It("should increase tail after eat", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        apple := [2]int{0, 3}

        newSnake := moveSnake(snake, apple, []int{right})
        Expect(len(newSnake)).To(Equal(len(snake) + 1))
    })

    // @TODO
    PIt("should prevent reverse moves", func() {

    })

    /*
        {1, 1, 1, 1, 0},
        {1, 0, 0, 0, 0},
        {1, 0, 0, 0, 0},
        {1, 1, 1, 1, 1},
        {0, 2, 0, 0, 0},
    */
    It("should not enter in closed zone", func() {
        snake := [][2]int{
            {0, 3}, {0, 2}, {0, 1}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4},
        }

        apple := [2]int{4, 1}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:1][0]).To(Equal(up))
    })

    /*
        {0, 0, 0, 1, 2},
        {0, 0, 0, 1, 1},
        {0, 0, 0, 0, 1},
        {0, 0, 0, 0, 1},
        {0, 0, 0, 0, 0},
    */
    It("should not eat apple if no free space", func() {
        snake := [][2]int{
            {3, 4}, {2, 4}, {1, 4}, {1, 3}, {0, 3},
        }

        apple := [2]int{0, 4}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:1][0]).To(Equal(left))
    })

    It("should select best path", func() {
        paths := []path{
            {
                []int{1, 1, 1, 1},
                80,
            },
            {
                []int{2, 1, 2, 1},
                90,
            },
            {
                []int{1, 2, 1, 0},
                50,
            },
        }

        Expect(getBestPath(paths).Score).To(Equal(float32(90)))
    })
})
