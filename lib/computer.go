package computer

const maxTick = 5
const width, height = 5, 5
const up, right, down, left = 0, 1, 2, 3
const block = 1;
const apple = 2;

func MoveSnake(snake [][2]int, moves []int) [][2]int {

    for _, move := range moves {
        snakeHead := GetSnakeHead(snake)
        nextPosition := GetAdjacentPosition(snakeHead, move)

        snake = snake[1:]
        snake = append(snake, nextPosition)
    }

    return snake
}

func GetAdjacentPosition(position [2]int, move int) [2]int {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    switch move {
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
    return GetSnakeHead(snake) == position
}

func IsOutsideBoundingBox(position [2]int, grid [width][height]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    var width = len(grid[0])
    var height = len(grid)

    if x > width || x < 0 || y > height || y < 0 {
        return true
    }

    return false
}

func getMoveScore(move int, snake [][2]int, apple [2]int) int {
    newSnake := MoveSnake(snake, []int{move})

    if isSnakeHeadAtPosition(newSnake, apple) {
        return 10
    }

    return 1
}

func IsEmptyCell(grid [width][height]int, position [2]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    if grid[x][y] != block {
        return true
    }

    return false
}

func GetPossibleMoves(grid [width][height]int, snake [][2]int) []int {
    var head = GetSnakeHead(snake)

    var possibleMoves []int
    for _, move := range []int{up, right, down, left} {
        var adjacentPosition = GetAdjacentPosition(head, move)

        if !IsOutsideBoundingBox(adjacentPosition, grid) && IsEmptyCell(grid, adjacentPosition) {
            possibleMoves = append(possibleMoves, move)
        }
	}

    return possibleMoves
}

func GetSnakeHead(snake [][2]int) [2]int {
    return snake[len(snake) - 1]
}

func getBestPath(paths [][]int, scores []int) []int {
    return paths[0] // @TODO
}

func GetPath(grid [width][height]int, snake [][2]int, apple [2]int) []int {
    var paths [][]int
    var scores []int

    var possibleMoves = GetPossibleMoves(grid, snake)

    for _, possibleMove := range possibleMoves {
        scores = append(scores, getMoveScore(possibleMove, snake, apple))
        paths = append(paths, []int{possibleMove})
    }

    for tick := 1; tick < maxTick; tick++ {
        for _, path := range paths {
            newSnake := MoveSnake(snake, path)

            for _, possibleMove := range GetPossibleMoves(grid, newSnake) {
                newPath := append(path, possibleMove)
                scores = append(scores, getMoveScore(possibleMove, newSnake, apple))
                paths = append(paths, newPath)
            }
        }
    }

    return getBestPath(paths, scores)
}
