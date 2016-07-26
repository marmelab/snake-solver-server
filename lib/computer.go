package lib

const maxTick = 5
const width, height = 5, 5
const up, right, down, left = 0, 1, 2, 3

func moveSnake(snake [][2]int, path []int) [][2]int {

    for direction := range path {
        snakeHead := getSnakeHead(snake)
        nextPosition := getAdjacentPosition(snakeHead, direction)

        snake = snake[1:]
        snake = append(snake, nextPosition)
    }

    return snake
}

func getAdjacentPosition(position [2]int, direction int) [2]int {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    switch direction {
    case up:
        return [2]int{x - 1, y}
    case right:
        return [2]int{x, y + 1}
    case down:
        return [2]int{x + 1, y}
    case left:
        return [2]int{x, y - 1}
    default:
        return [2]int{x, y + 1} // @FIXME: return false
    }
}

func isSnakeHeadAtPosition(snake [][2]int, position [2]int) bool {
    return getSnakeHead(snake) == position
}

func isOutsideBoundingBox(position [2]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    if x > width || x < 0 || y > height || y < 0 {
        return true
    }

    return false
}

func getMoveScore(snake [][2]int, apple [2]int) int {
    if isSnakeHeadAtPosition(snake, apple) {
        return 10
    }

    return 1
}

func getPossibleDirections(grid [width][height]int, snake [][2]int) []int {
    var head = getSnakeHead(snake)

    var possibleDirections []int
    for _, direction := range []int{up, right, down, left} {
        var adjacentPosition = getAdjacentPosition(head, direction)

        if !isOutsideBoundingBox(adjacentPosition) {
            possibleDirections = append(possibleDirections, direction)
        }
	}

    return possibleDirections
}

func getSnakeHead(snake [][2]int) [2]int {
    return snake[len(snake) - 1]
}

func getBestPath(paths [][]int, scores []int) []int {
    return paths[0] // @TODO
}

func GetPath(grid [width][height]int, snake [][2]int, apple [2]int) []int {
    var paths [][]int
    var scores []int

    var possibleDirections = getPossibleDirections(grid, snake)

    for _, possibleDirection := range possibleDirections {
        paths = append(paths, []int{possibleDirection})
    }

    for tick := 1; tick < maxTick; tick++ {
        for index, path := range paths {
            newSnake := moveSnake(snake, path)

            for _, possibleDirection := range getPossibleDirections(grid, newSnake) {
                paths[index] = append(paths[index], possibleDirection)
            }
        }
    }

    return getBestPath(paths, scores)
}
